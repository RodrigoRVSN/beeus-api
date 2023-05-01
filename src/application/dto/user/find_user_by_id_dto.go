package userDTO

import "time"

type FindUserByIdOutputDTO struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	AvatarUrl string    `json:"avatar_url"`
	Points    int       `json:"points"`
	CreatedAt time.Time `json:"created_at"`
}
