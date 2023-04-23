package userRepository

import (
	"github.com/rodrigoRVSN/beeus-api/src/application/dto"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepositoryImplementation interface {
	CreateUser(payload *entity.User) error
	SignInUser(payload dto.SignInInputDTO) (*dto.SignInOutputDTO, error)
	FindAllUsers(user *entity.User) ([]entity.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepositoryImplementation {
	return &UserRepository{DB: db}
}
