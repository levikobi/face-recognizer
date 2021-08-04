package models

import (
	"encoding/json"
	pb "grpc-schemas/golang/face-recognition/protos"
	"io"
)

type Match struct {
	Person Person
	Accuracy float64
}

func ToProto(match Match, proto *pb.Person) {
	proto.Features = match.Person.Features
	proto.Name = match.Person.Name
}

func ToJSON(i interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(i)
}

func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
