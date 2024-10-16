package utils

import "github.com/go-playground/validator"

type Validator struct {
	validate *validator.Validate
}

var val *Validator

func InitValidator() {
	val = &Validator{}
	val.validate = validator.New()
}

func Validate(v interface{}) error {
	return val.validate.Struct(v)
}
