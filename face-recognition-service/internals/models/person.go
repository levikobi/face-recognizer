package models

import (
	"encoding/json"
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
