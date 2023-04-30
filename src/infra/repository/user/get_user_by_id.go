package userRepository

import (
	"github.com/rodrigoRVSN/beeus-api/src/application/dto"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (r *UserRepository) GetUserById(userId uint) (*dto.FindUserByIdOutputDTO, error) {
	user := entity.User{}

	response := r.DB.First(&user, userId)

	if response.Error != nil {
		return nil, response.Error
	}

	return &dto.FindUserByIdOutputDTO{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		AvatarUrl: user.AvatarUrl,
		Points:    user.Points,
		CreatedAt: user.CreatedAt,
	}, nil
}
