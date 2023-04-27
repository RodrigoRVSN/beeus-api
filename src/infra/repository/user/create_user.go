package userRepository

import (
	"errors"

	"github.com/rodrigoRVSN/beeus-api/src/application/dto"
	hash "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
	"gorm.io/gorm"
)

func (r *UserRepository) CheckIfUserAlreadyExists(payload dto.SignUpInputDTO, email string) error {
	if err := r.DB.First(payload, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return errors.New("houve um erro ao cadastrar o email")
}

func (r *UserRepository) CreateUser(payload dto.SignUpInputDTO) error {
	payload.Password = hash.HashPassword(payload.Password)

	response := r.DB.Create(payload)

	if response.Error != nil {
		return response.Error
	}

	return nil
}
