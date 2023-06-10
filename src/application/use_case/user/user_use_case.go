package userUseCase

import (
	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
	"github.com/rodrigoRVSN/beeus-api/src/domain/gateway"
)

type UserUseCaseInterface interface {
	CreateUser(payload userDTO.SignUpInputDTO) error
	SignInUser(payload userDTO.SignInInputDTO) (*userDTO.SignInOutputDTO, error)
	Me(userId uint) (*userDTO.FindUserByIdOutputDTO, error)
	GetRankedUsers() (*[]userDTO.UserWithoutPasswordDTO, error)
}

type UserUseCase struct {
	UserGateway gateway.UserGateway
}

func NewUserUseCase(userGateway gateway.UserGateway) UserUseCaseInterface {
	return &UserUseCase{
		UserGateway: userGateway,
	}
}
