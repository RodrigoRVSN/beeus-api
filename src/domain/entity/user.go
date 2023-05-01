package entity

import (
	"time"
)

type User struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	AvatarUrl string    `json:"avatar_url" gorm:"default:https://avatarfiles.alphacoders.com/260/thumb-260445.jpg"`
	Password  string    `json:"password" validate:"required"`
	Points    int       `json:"points" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at"`
}
