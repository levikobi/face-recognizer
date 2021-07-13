package handlers

import (
	"container/heap"
	"context"
	"encoding/json"
	"face-recognition-service/internals/models"
	"fmt"
	"gonum.org/v1/gonum/floats"
	"log"
	"math"
	"net/http"
)

func (this *Api) FindPerson(res http.ResponseWriter, req *http.Request) {
	log.Println("Handle FindPerson request")

	person := &models.Person{}

	// TODO: use middleware here to deserialize and validate
	if err := person.FromJson(req.Body); err != nil {
		log.Println("[ERROR] deserializing features", err)
		http.Error(res, "Error reading features", http.StatusBadRequest)
		return
	}
	person.SquareOfMagnitude = math.Sqrt(floats.Dot(person.Features, person.Features))

	persons, err := this.persons.Find(context.Background(), json.Encoder{})
	if err != nil {
		log.Println("[ERROR] fetching persons from DB", err)
		http.Error(res, "Error fetching persons", http.StatusInternalServerError)
		return
	}

	topMatches := this.getTopMatches(person, persons)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusFound)

	for topMatches.Len() > 0 {
		m := (heap.Pop(topMatches)).(models.Match)
		models.ToJSON(m.Person, res)
	}
}

func (this *Api) getTopMatches(receivedPerson *models.Person, persons *[]models.Person) *models.Heap {
	topMatches := &models.Heap{}
	heap.Init(topMatches)

	featuresString := fmt.Sprint(receivedPerson)
	for _, currPerson := range *persons {
		angle := this.calculateAngle(receivedPerson, currPerson, featuresString)
		this.updateTopMatches(currPerson, angle, topMatches)
	}
	return topMatches
}

func (this *Api) calculateAngle(receivedPerson *models.Person, currPerson models.Person, featuresString string) float64 {
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

func (this *Api) checkConnection(f1, f2 string) (float64, bool, error) {
	angleByFeatures, exists, err := this.angles.Get(context.Background(), f1)
	if err != nil || !exists {
		return 0, false, err
	}
	value, exists := angleByFeatures[f2]
	return value, exists, nil
}

func (this *Api) updateTopMatches(currPerson models.Person, angle float64, topMatches *models.Heap) {
	match := &models.Match{
		Person:   currPerson,
		Accuracy: angle,
	}
	heap.Push(topMatches, *match)
	if topMatches.Len() > this.config.NumberOfBestMatches {
		heap.Pop(topMatches)
	}
}
