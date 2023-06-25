package utils

import "github.com/go-playground/validator"

// Validator is the interface that wraps the Validate function.
type StructValidator interface {
	Validate(i interface{}) error
}

type structValidator struct {
	validator *validator.Validate
}

func NewStructValidator() StructValidator {
	valid := validator.New()
	return &structValidator{validator: valid}
}

func (v structValidator) Validate(i interface{}) error {
	err := v.validator.Struct(i)
	if err != nil {
		return err
	}
	return nil
}
