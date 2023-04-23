package authService

import (
	database "github.com/rodrigoRVSN/beeus-api/src/config"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User

	if err := database.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
