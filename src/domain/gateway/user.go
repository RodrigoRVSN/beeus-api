package gateway

import (
	"github.com/rodrigoRVSN/beeus-api/src/application/dto"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

type UserGateway interface {
	CreateUser(payload *entity.User) error
	SignInUser(payload dto.SignInInputDTO) (*dto.SignInOutputDTO, error)
	FindAllUsers(user *entity.User) ([]entity.User, error)
}
