package domain

import "github.com/go-playground/validator"

var Validator *validator.Validate

func ValidatorSingletonClient() {
	Validator = validator.New()
}
