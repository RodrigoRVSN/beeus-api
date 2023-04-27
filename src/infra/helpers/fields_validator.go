package helpers

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type ValidationError struct {
	Errors []*ErrorResponse
}

var validate = validator.New()

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s", e.FailedField, e.Tag)
}

func (e *ValidationError) Error() string {
	errorStr := "Validation failed:"
	for _, err := range e.Errors {
		errorStr += fmt.Sprintf(" %s", err.Error())
	}
	return errorStr
}

func messageByTag(tag string) string {
	switch tag {
	case "required":
		return "Esse campo é obrigatório"
	case "email":
		return "Email inválido"
	}
	return tag
}

func ValidateStruct(user interface{}) *ValidationError {
	var errors []*ErrorResponse
	err := validate.Struct(user)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = messageByTag(err.Tag())
			errors = append(errors, &element)
		}
	}

	if len(errors) > 0 {
		return &ValidationError{Errors: errors}
	}

	return nil
}
