package authService

import (
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
	database "github.com/rodrigoRVSN/beeus-api/src/infra/config"
)

func FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User

	if err := database.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
