package main

import (
	"context"
	"face-recognition-service/internals/data/angles/redis_client"
	"face-recognition-service/internals/data/persons/mongo_client"
	"face-recognition-service/internals/handlers"
	"face-recognition-service/internals/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	config, err := models.ParseConfig()
	if err != nil { log.Fatal(err) }

	api := handlers.NewApi(config,
		mongo_client.Connect(config),
		redis_client.Connect(config))

	sm := mux.NewRouter()

	addPersonRouter := sm.Methods(http.MethodPost).Subrouter()
	addPersonRouter.HandleFunc("/api/person/", api.AddPerson)

	findPersonRouter := sm.Methods(http.MethodGet).Subrouter()
	findPersonRouter.HandleFunc("/api/person/", api.FindPerson)

	server := &http.Server{
		Addr:              ":3000",
		Handler:           sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		log.Println("face-recognition-service is running")
		if err = server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	log.Println("Received terminate, graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	server.Shutdown(timeoutContext)
}
