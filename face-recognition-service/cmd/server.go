package cmd

import (
	"context"
	"face-recognition-service/internals/business"
	"face-recognition-service/internals/models"
	pb "grpc-schemas/golang/face-recognition/protos"
)

type Server struct {
	pb.UnimplementedFaceRecognitionServer

	service *business.FaceRecognitionService
}

func NewServer(service *business.FaceRecognitionService) *Server {
	return &Server{service: service}
}

func (this *Server) AddPerson(ctx context.Context, in *pb.AddPersonRequest) (*pb.AddPersonResponse, error) {
	var person models.Person
	models.FromProto(&person, in.Person)

	if err := this.service.AddPerson(&person); err != nil {
		return &pb.AddPersonResponse{Status: pb.AddPersonResponse_FAILURE} , err
	}

	return &pb.AddPersonResponse{Status: pb.AddPersonResponse_SUCCESS}, nil
}

func (this *Server) FindPerson(ctx context.Context, in *pb.FindPersonRequest) (*pb.FindPersonResponse, error) {
	var person models.Person
	person.Features = in.Features

	matches, err := this.service.FindPerson(&person)
	if err != nil {
		return &pb.FindPersonResponse{Persons: nil}, err
	}

	var persons []*pb.Person
	for i := range matches {
		persons = append(persons, &pb.Person{
			Features: matches[i].Features,
			Name:     matches[i].Name,
		})
	}
	return &pb.FindPersonResponse{Persons: persons}, nil
}