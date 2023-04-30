package userUseCase

import "github.com/rodrigoRVSN/beeus-api/src/application/dto"

func (uc *UserUseCase) Me(userId uint) (*dto.FindUserByIdOutputDTO, error) {
	user, err := uc.UserGateway.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}
