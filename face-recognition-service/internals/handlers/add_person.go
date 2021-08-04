package handlers

import (
	"context"
	"encoding/json"
	"face-recognition-service/internals/models"
	"gonum.org/v1/gonum/floats"
	"log"
	"math"
	"net/http"
)

func (this *Api) AddPerson(res http.ResponseWriter, req *http.Request) {
	log.Println("Handle AddPerson request")




	numOfPeople, _ := this.persons.Count(context.Background(), json.Encoder{})
	if numOfPeople >= this.config.MaxAmountOfPersons {
		http.Error(res, "Cannot add more people", http.StatusMethodNotAllowed)
		return
	}

	person := &models.Person{}

	// TODO: use middleware here to deserialize and validate
	if err := person.FromJson(req.Body); err != nil {
		log.Println("[ERROR] deserializing person", err)
		http.Error(res, "Error reading person", http.StatusBadRequest)
		return
	}

	person.SquareOfMagnitude = math.Sqrt(floats.Dot(person.Features, person.Features))

	if err := this.persons.InsertOne(context.Background(), person); err != nil {
		log.Println("[ERROR] inserting person to db", err)
		http.Error(res, "Error inserting person to db", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
}
