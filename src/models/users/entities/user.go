package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (b *User) BeforeCreate(tx *gorm.DB) (err error) {
	b.Uuid = uuid.New()
	return
}

type User struct {
	Uuid      uuid.UUID `json:"uuid" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" gorm:"not null" validate:"required"`
	AvatarUrl string    `json:"avatar_url" gorm:"default:https://avatarfiles.alphacoders.com/260/thumb-260445.jpg"`
	Points    int       `json:"points" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at"`
}
