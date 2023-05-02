package userUseCase

import (
	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
)

func (uc *UserUseCase) SignInUser(payload userDTO.SignInInputDTO) (*userDTO.SignInOutputDTO, error) {
	user, err := uc.UserGateway.SignInUser(payload)

	if err != nil {
		return nil, err
	}

	return user, nil
}
