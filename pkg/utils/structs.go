package utils

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) error {
	return validator.New().Struct(s)
}

func StructToJSON(s interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(s)
}

func JSONToStuct(s interface{}, r io.Reader) error  {
	return json.NewDecoder(r).Decode(&s)
}