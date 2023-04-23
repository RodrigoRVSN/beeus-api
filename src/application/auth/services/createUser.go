package authService

import (
	"github.com/rodrigoRVSN/beeus-api/src/application/users/entities"
	database "github.com/rodrigoRVSN/beeus-api/src/infra/config"
	hash "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
	fieldsValidator "github.com/rodrigoRVSN/beeus-api/src/infra/helpers/validator"
)

func CreateUser(payload *entities.User) error {
	errors := fieldsValidator.ValidateStruct(payload)

	if errors != nil {
		return errors
	}

	payload.Password = hash.HashPassword(payload.Password)

	database.DB.Create(&payload)

	return nil
}
