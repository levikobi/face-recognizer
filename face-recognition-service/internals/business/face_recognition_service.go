package business

import (
	"context"
	"encoding/json"
	"errors"
	"face-recognition-service/internals/data/angles"
	"face-recognition-service/internals/data/persons"
	"face-recognition-service/internals/models"
	"gonum.org/v1/gonum/floats"
	"math"
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
		return errors.New("cannot add more people")
	}

	person.SquareOfMagnitude = math.Sqrt(floats.Dot(person.Features, person.Features))

	if err := this.persons.InsertOne(context.Background(), person); err != nil {
		return errors.New("failed to insert person to db")
	}

	return nil
}

func (this *FaceRecognitionService) FindPerson(person *models.Person) ([]models.Person, error) {
	panic("not implemented")
}