package userRepository

import (
	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (r *UserRepository) GetUserById(userId uint) (*userDTO.FindUserByIdOutputDTO, error) {
	user := entity.User{}

	if err := r.DB.First(&user, userId).Error; err != nil {
		return nil, err
	}

	return &userDTO.FindUserByIdOutputDTO{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		AvatarUrl: user.AvatarUrl,
		Points:    user.Points,
		CreatedAt: user.CreatedAt,
	}, nil
}
