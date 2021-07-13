package handlers

import (
	"face-recognition-service/internals/data/persons"
	"face-recognition-service/internals/data/persons/mongo_client"
	"face-recognition-service/internals/models"
	"log"
	"testing"
)

var api *Api
var personsCollection persons.Repository

func TestMain(m *testing.M) {
	config, err := models.ParseConfig()
	if err != nil { log.Fatal(err) }

	personsCollection = mongo_client.Connect(config)
	api = NewApi(config, personsCollection, nil)

	m.Run()
}
