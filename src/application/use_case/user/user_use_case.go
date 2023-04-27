package userUseCase

import (
	"github.com/rodrigoRVSN/beeus-api/src/domain/gateway"
)

type CreateUserUseCase struct {
	UserGateway gateway.UserGateway
}

func NewUserUseCase(chatGateway gateway.UserGateway) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserGateway: chatGateway,
	}
}
