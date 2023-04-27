package userRepository

import (
	"github.com/rodrigoRVSN/beeus-api/src/domain/gateway"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) gateway.UserGateway {
	return &UserRepository{DB: db}
}
