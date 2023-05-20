package userRepository

import (
	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
	hash "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
)

func (r *UserRepository) CreateUser(payload userDTO.SignUpInputDTO) error {
	payload.Password = hash.HashPassword(payload.Password)

	user := entity.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}

	if err := r.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
