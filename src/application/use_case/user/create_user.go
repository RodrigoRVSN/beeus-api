package userUseCase

import (
	"github.com/rodrigoRVSN/beeus-api/src/application/dto"
)

func (uc *CreateUserUseCase) CreateUser(payload dto.SignUpInputDTO) error {
	userWasFound := uc.UserGateway.CheckIfUserAlreadyExists(payload, payload.Email)

	if userWasFound != nil {
		return userWasFound
	}

	err := uc.UserGateway.CreateUser(payload)

	if err != nil {
		return err
	}

	return nil
}
