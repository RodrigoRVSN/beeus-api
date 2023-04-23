package authService

import (
	database "github.com/rodrigoRVSN/beeus-api/src/config"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
	hash "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
	fieldsValidator "github.com/rodrigoRVSN/beeus-api/src/infra/helpers/validator"
)

func CreateUser(payload *entity.User) error {
	errors := fieldsValidator.ValidateStruct(payload)

	if errors != nil {
		return errors
	}

	payload.Password = hash.HashPassword(payload.Password)

	database.DB.Create(&payload)

	return nil
}
