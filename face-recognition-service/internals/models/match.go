package models

import (
	"encoding/json"
	"io"
)

type Match struct {
	Person Person
	Accuracy float64
}

func ToJSON(i interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(i)
}

func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
