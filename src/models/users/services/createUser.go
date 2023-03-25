package userService

import (
	database "github.com/rodrigoRVSN/beeus-api/src/infra/config"
	hash "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
	fieldsValidator "github.com/rodrigoRVSN/beeus-api/src/infra/helpers/validator"
	"github.com/rodrigoRVSN/beeus-api/src/models/users/entities"
)

func CreateUser(payload *entities.User) error {
	errors := fieldsValidator.ValidateStruct(payload)

	if errors != nil {
		return errors
	}

	payload.Password = hash.HashPassword(payload.Password)

	database.DB.Db.Create(&payload)

	return nil
}
