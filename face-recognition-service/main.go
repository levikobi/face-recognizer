package main

import (
	"face-recognition-service/cmd"
	"face-recognition-service/internals/business"
	"face-recognition-service/internals/data/angles/redis_client"
	"face-recognition-service/internals/data/persons/mongo_client"
	"face-recognition-service/internals/models"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-schemas/golang/face-recognition/protos"
	"log"
	"net"
	"os"
)

func main() {
	config, err := models.ParseConfig()
	if err != nil { log.Fatal(err) }

	service := business.New(config,
		mongo_client.Connect(config),
		redis_client.Connect(config))

	gs := grpc.NewServer()
	c := cmd.NewServer(service)

	pb.RegisterFaceRecognitionServer(gs, c)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		os.Exit(1)
	}

	// listen for requests
	go func() {
		gs.Serve(l)
	}()

	test_client()

}

//func main() {
//	config, err := models.ParseConfig()
//	if err != nil { log.Fatal(err) }
//
//	api := handlers.NewApi(config,
//		mongo_client.Connect(config),
//		redis_client.Connect(config))
//
//	sm := mux.NewRouter()
//
//	addPersonRouter := sm.Methods(http.MethodPost).Subrouter()
//	addPersonRouter.HandleFunc("/api/person/", api.AddPerson)
//
//	findPersonRouter := sm.Methods(http.MethodGet).Subrouter()
//	findPersonRouter.HandleFunc("/api/person/", api.FindPerson)
//
//	server := &http.Server{
//		Addr:              ":3000",
//		Handler:           sm,
//		IdleTimeout: 120 * time.Second,
//		ReadTimeout: 1 * time.Second,
//		WriteTimeout: 1 * time.Second,
//	}
//
//	go func() {
//		log.Println("face-recognition-service is running")
//		if err = server.ListenAndServe(); err != nil {
//			log.Fatal(err)
//		}
//	}()
//
//	sigChan := make(chan os.Signal)
//	signal.Notify(sigChan, os.Interrupt)
//	signal.Notify(sigChan, os.Kill)
//
//	sig := <- sigChan
//	log.Println("Received terminate, graceful shutdown", sig)
//
//	timeoutContext, _ := context.WithTimeout(context.Background(), 30 * time.Second)
//	server.Shutdown(timeoutContext)
//}
