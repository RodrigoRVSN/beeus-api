package CreateUserValidation

import (
	"github.com/rodrigoRVSN/beeus-api/src/models/users/entities"
	"gopkg.in/go-playground/validator.v9"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func messageByTag(tag string) string {
	switch tag {
	case "required":
		return "Esse campo é obrigatório"
	case "email":
		return "Email inválido"
	}
	return tag
}

func ValidateStruct(user entities.User) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = messageByTag(err.Tag())
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
