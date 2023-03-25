package userService

import (
	database "github.com/rodrigoRVSN/beeus-api/src/infra/config"
	hash "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
	"github.com/rodrigoRVSN/beeus-api/src/models/users/entities"
	CreateUserValidation "github.com/rodrigoRVSN/beeus-api/src/models/users/validations"
)

func CreateUserService(user *entities.User) error {
	errors := CreateUserValidation.ValidateStruct(*user)

	if errors != nil {
		return errors
	}

	user.Password = hash.HashPassword(user.Password)

	database.DB.Db.Create(&user)

	return nil
}
