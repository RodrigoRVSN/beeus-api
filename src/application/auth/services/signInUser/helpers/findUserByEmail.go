package authService

import (
	"github.com/rodrigoRVSN/beeus-api/src/application/users/entities"
	database "github.com/rodrigoRVSN/beeus-api/src/infra/config"
)

func FindUserByEmail(email string) (*entities.User, error) {
	var user entities.User

	if err := database.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
