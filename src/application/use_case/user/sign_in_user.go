package userUseCase

import (
	"github.com/rodrigoRVSN/beeus-api/src/application/dto"
)

func (uc *UserUseCase) SignInUser(payload dto.SignInInputDTO) (*dto.SignInOutputDTO, error) {
	user, err := uc.UserGateway.SignInUser(payload)

	if err != nil {
		return nil, err
	}

	return user, nil
}
