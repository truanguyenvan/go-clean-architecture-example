package structure

import "github.com/go-playground/validator"

// Validator is the interface that wraps the Validate function.
type Validator interface {
	Validate(i interface{}) error
}

type _validator struct {
	validator *validator.Validate
}

func NewValidator() Validator {
	valid := validator.New()
	return &_validator{validator: valid}
}

func (v _validator) Validate(i interface{}) error {
	err := v.validator.Struct(i)
	if err != nil {
		return err
	}
	return nil
}
