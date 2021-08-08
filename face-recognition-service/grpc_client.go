package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-schemas/golang/face-recognition/protos"
	"log"
)

var (
	serverAddr         = flag.String("server_addr", "localhost:3000", "The server address in the format of host:port")
)

func test_client() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewFaceRecognitionClient(conn)

	resp, err := client.AddPerson(context.Background(), &pb.AddPersonRequest{Person: &pb.Person{
		Features: []float64{0.1, 0.4, 0.2, 0.9, 0.7},
		Name:     "Kobi",
	}})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)
}