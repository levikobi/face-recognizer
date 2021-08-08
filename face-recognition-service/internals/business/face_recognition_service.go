package business

import (
	"container/heap"
	"context"
	"encoding/json"
	"errors"
	"face-recognition-service/internals/data/angles"
	"face-recognition-service/internals/data/persons"
	"face-recognition-service/internals/models"
	"fmt"
	"gonum.org/v1/gonum/floats"
	"math"
	"sync"
)

type FaceRecognitionService struct {
	config models.Config
	persons persons.Repository
	angles angles.Repository
}

func New(config models.Config, persons persons.Repository, angles angles.Repository) *FaceRecognitionService {
	return &FaceRecognitionService{config, persons, angles}
}

func (this *FaceRecognitionService) AddPerson(person *models.Person) error {
	numOfPeople, _ := this.persons.Count(context.Background(), json.Encoder{})
	if numOfPeople >= this.config.MaxAmountOfPersons {
		return errors.New("cannot add more people, reached limit")
	}

	person.SquareOfMagnitude = math.Sqrt(floats.Dot(person.Features, person.Features))

	if err := this.persons.InsertOne(context.Background(), person); err != nil {
		return err
	}

	return nil
}

func (this *FaceRecognitionService) FindPerson(person *models.Person) ([]models.Person, error) {
	person.SquareOfMagnitude = math.Sqrt(floats.Dot(person.Features, person.Features))

	allMatches, err := this.persons.Find(context.Background(), json.Encoder{})
	if err != nil {
		return nil, err
	}

	topMatches := this.getTopMatches(person, allMatches)

	var bestMatches []models.Person
	for topMatches.Len() > 0 {
		m := (heap.Pop(topMatches)).(models.Match)
		bestMatches = append(bestMatches, m.Person)
	}

	return bestMatches, nil
}

func (this *FaceRecognitionService) getTopMatches(receivedPerson *models.Person, persons *[]models.Person) *models.Heap {
	topMatches := &models.Heap{}
	heap.Init(topMatches)

	var wg sync.WaitGroup
	var mu = &sync.Mutex{}

	numJobs := len(*persons)
	jobs := make(chan models.Person, numJobs)
	wg.Add(numJobs)

	featuresString := fmt.Sprint(receivedPerson)

	for w := 1; w <= this.config.NumberOfWorkers; w++ {
		go func (w int) {
			for currPerson := range jobs {
				angle := this.calculateAngle(receivedPerson, currPerson, featuresString)

				mu.Lock()
				this.updateTopMatches(currPerson, angle, topMatches)
				mu.Unlock()

				wg.Done()
			}
		}(w)
	}

	for _, currPerson := range *persons {
		jobs <- currPerson
	}

	close(jobs)
	wg.Wait()

	return topMatches
}

func (this *FaceRecognitionService) calculateAngle(receivedPerson *models.Person, currPerson models.Person, featuresString string) float64 {
	currFeaturesString := fmt.Sprint(currPerson.Features)
	angle, found, err := this.checkConnection(featuresString, currFeaturesString)

	if !found || err != nil {
		angle, found, err = this.checkConnection(currFeaturesString, featuresString)
		if !found {
			angle = math.Acos(floats.Dot(receivedPerson.Features, currPerson.Features) / (receivedPerson.SquareOfMagnitude * currPerson.SquareOfMagnitude))

			angleByFeatures, exists, _ := this.angles.Get(context.Background(), featuresString)
			if !exists {
				angleByFeatures = make(map[string]float64)
			}
			angleByFeatures[currFeaturesString] = angle
			this.angles.Set(context.Background(), featuresString, &angleByFeatures)
		}
	}
	if found {
		fmt.Println("cache hit with person: ", currFeaturesString)
	}
	return angle
}

func (this *FaceRecognitionService) checkConnection(f1, f2 string) (float64, bool, error) {
	angleByFeatures, exists, err := this.angles.Get(context.Background(), f1)
	if err != nil || !exists {
		return 0, false, err
	}
	value, exists := angleByFeatures[f2]
	return value, exists, nil
}

func (this *FaceRecognitionService) updateTopMatches(currPerson models.Person, angle float64, topMatches *models.Heap) {
	match := &models.Match{
		Person:   currPerson,
		Accuracy: angle,
	}
	heap.Push(topMatches, *match)
	if topMatches.Len() > this.config.NumberOfBestMatches {
		heap.Pop(topMatches)
	}
}

