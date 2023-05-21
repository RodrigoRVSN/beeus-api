package userUseCase

import (
	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
	"github.com/rodrigoRVSN/beeus-api/src/domain/gateway"
)

type UserControllerInterface interface {
	CreateUser(payload userDTO.SignUpInputDTO) error
	SignInUser(payload userDTO.SignInInputDTO) (*userDTO.SignInOutputDTO, error)
	FindAllUsers(user *entity.User) ([]entity.User, error)
	Me(userId uint) (*userDTO.FindUserByIdOutputDTO, error)
}

type UserUseCase struct {
	UserGateway gateway.UserGateway
}

func NewUserUseCase(userGateway gateway.UserGateway) UserControllerInterface {
	return &UserUseCase{
		UserGateway: userGateway,
	}
}
