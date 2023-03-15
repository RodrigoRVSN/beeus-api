package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	AvatarUrl string    `json:"avatar_url"`
	Password  string    `json:"password" gorm:"default:"`
	Points    int       `json:"points" gorm:"default:https://avatarfiles.alphacoders.com/260/thumb-260445.jpg"`
	CreatedAt time.Time `json:"created_at" gorm:"default:0"`
}
