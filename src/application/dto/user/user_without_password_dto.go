package userDTO

import (
	"time"

	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

type UserWithoutPasswordDTO struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	AvatarUrl string    `json:"avatar_url"`
	Points    int       `json:"points"`
	CreatedAt time.Time `json:"created_at"`
}

func MapUserWithoutPasswordToDTO(user entity.User) UserWithoutPasswordDTO {
	return UserWithoutPasswordDTO{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		AvatarUrl: user.AvatarUrl,
		Points:    user.Points,
		CreatedAt: user.CreatedAt,
	}
}
