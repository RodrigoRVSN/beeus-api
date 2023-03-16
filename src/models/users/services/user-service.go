package userService

import (
	database "github.com/rodrigoRVSN/beeus-api/src/infra/config"
	"github.com/rodrigoRVSN/beeus-api/src/models/users/entities"
	CreateUserValidation "github.com/rodrigoRVSN/beeus-api/src/models/users/validations"
)

func CreateUserService(user *entities.User) error {
	errors := CreateUserValidation.ValidateStruct(*user)

	if errors != nil {
		return errors
	}

	database.DB.Db.Create(&user)

	return nil
}
