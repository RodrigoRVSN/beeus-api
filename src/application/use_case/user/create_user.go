package userUseCase

import (
	"github.com/rodrigoRVSN/beeus-api/src/application/dto"
)

func (uc *UserUseCase) CreateUser(payload dto.SignUpInputDTO) error {
	userWasFound := uc.UserGateway.CheckIfUserAlreadyExists(payload.Email)

	if userWasFound != nil {
		return userWasFound
	}

	err := uc.UserGateway.CreateUser(payload)

	if err != nil {
		return err
	}

	return nil
}
