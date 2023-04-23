package userRepository

import (
	"fmt"

	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
	hash "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
)

func (r *UserRepository) CreateUser(payload *entity.User) error {
	payload.Password = hash.HashPassword(payload.Password)

	if userFound := r.DB.First(&payload, "email = ?", payload.Email); userFound != nil {
		return fmt.Errorf("houve um erro ao cadastrar o email")
	}

	response := r.DB.Create(payload)

	if response.Error != nil {
		return response.Error
	}

	return nil
}
