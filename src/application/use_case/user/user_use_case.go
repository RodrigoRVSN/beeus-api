package userUseCase

import (
	"github.com/rodrigoRVSN/beeus-api/src/domain/gateway"
)

type UserUseCase struct {
	UserGateway gateway.UserGateway
}

func NewUserUseCase(userGateway gateway.UserGateway) *UserUseCase {
	return &UserUseCase{
		UserGateway: userGateway,
	}
}
