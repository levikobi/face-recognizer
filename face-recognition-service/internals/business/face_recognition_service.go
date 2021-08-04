package business

import (
	"face-recognition-service/internals/data/angles"
	"face-recognition-service/internals/data/persons"
	"face-recognition-service/internals/models"
)

type FaceRecognitionService struct {
	persons persons.Repository
	angles angles.Repository
}

func (this *FaceRecognitionService) AddPerson(person *models.Person) error {
	panic("not implemented")
}

func (this *FaceRecognitionService) FindPerson(person *models.Person) ([]models.Person, error) {
	panic("not implemented")
}