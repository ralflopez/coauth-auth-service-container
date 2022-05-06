package utils

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) error {
	return validator.New().Struct(s)
}

func StructToJSON(w io.Writer, s interface{}) error {
	return json.NewEncoder(w).Encode(s)
}

func JSONToStuct(r io.Reader, s interface{}) error  {
	return json.NewDecoder(r).Decode(&s)
}