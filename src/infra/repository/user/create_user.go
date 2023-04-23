package userRepository

import (
	"errors"
	"fmt"

	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
	hash "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
	"gorm.io/gorm"
)

func (r *UserRepository) checkUserAlreadyExists(payload *entity.User, email string) error {
	if err := r.DB.First(payload, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return fmt.Errorf("houve um erro ao cadastrar o email")
}

func (r *UserRepository) CreateUser(payload *entity.User) error {
	payload.Password = hash.HashPassword(payload.Password)

	if err := r.checkUserAlreadyExists(payload, payload.Email); err != nil {
		return err
	}

	response := r.DB.Create(payload)

	if response.Error != nil {
		return response.Error
	}

	return nil
}
