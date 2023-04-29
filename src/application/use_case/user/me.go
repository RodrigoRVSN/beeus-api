package userUseCase

import (
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (uc *CreateUserUseCase) Me(userId uint) (*entity.User, error) {
	user, err := uc.UserGateway.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}
