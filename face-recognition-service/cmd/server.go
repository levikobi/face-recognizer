package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-schemas/golang/face-recognition/protos"
)

func main(){
	fmt.Println("Hello world")
}

type Server struct {

}

func (this *Server) AddPerson(ctx context.Context, in *pb.AddPersonRequest, opts ...grpc.CallOption) (*pb.AddPersonResponse, error) {

	panic("not implemented")
}

func (this *Server) FindPerson(ctx context.Context, in *pb.FindPersonRequest, opts ...grpc.CallOption) (*pb.FindPersonResponse, error) {
	panic("not implemented")
}