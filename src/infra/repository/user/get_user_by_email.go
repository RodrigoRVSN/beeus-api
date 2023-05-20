package userRepository

import (
	"errors"

	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (r *UserRepository) GetUserByEmail(email string) error {
	user := entity.User{}

	result := r.DB.Where("email = ?", email).First(&user)

	if result.RowsAffected > 0 {
		return errors.New("o email já está em uso")
	}

	return nil
}
