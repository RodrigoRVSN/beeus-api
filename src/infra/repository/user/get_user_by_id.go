package userRepository

import (
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (r *UserRepository) GetUserById(userId uint) (*entity.User, error) {
	user := entity.User{}

	response := r.DB.First(&user, userId)

	if response.Error != nil {
		return nil, response.Error
	}

	return &user, nil
}
