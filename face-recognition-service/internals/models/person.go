package models

import (
	"encoding/json"
	pb "grpc-schemas/golang/face-recognition/protos"
	"io"
)

type Person struct {
	Name string `json:"name"`
	Features []float64 `json:"features"`
	SquareOfMagnitude float64 `json:"-"`
}

func (this *Person) FromJson(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(this)
}

func FromProto(person *Person, proto *pb.Person) {
	person.Features = person.Features
	person.Name = person.Name
}