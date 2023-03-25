package userService

import (
	database "github.com/rodrigoRVSN/beeus-api/src/infra/config"
	hash "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
	fieldsValidator "github.com/rodrigoRVSN/beeus-api/src/infra/helpers/validator"
	"github.com/rodrigoRVSN/beeus-api/src/models/users/entities"
)

func CreateUserService(user *entities.User) error {
	errors := fieldsValidator.ValidateStruct(user)

	if errors != nil {
		return errors
	}

	user.Password = hash.HashPassword(user.Password)

	database.DB.Db.Create(&user)

	return nil
}
