package userUseCase

import (
	"github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
)

func (uc *UserUseCase) CreateUser(payload userDTO.SignUpInputDTO) error {
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
