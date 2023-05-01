package entity

import "time"

type Documentation struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"email" validate:"required,email"`
	CreatedAt time.Time `json:"created_at" gorm:"default:0"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:0"`
	UserID    uint
	User      User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
