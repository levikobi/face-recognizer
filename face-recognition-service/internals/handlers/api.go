package handlers

import (
	"face-recognition-service/internals/data/angles"
	"face-recognition-service/internals/data/persons"
	"face-recognition-service/internals/models"
)

type Api struct {
	config models.Config
	persons persons.Repository
	angles angles.Repository
}

func NewApi(config models.Config, persons persons.Repository, angles angles.Repository) *Api {
	return &Api{config, persons, angles}
}
