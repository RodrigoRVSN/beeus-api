package userUseCase

import "github.com/rodrigoRVSN/beeus-api/src/domain/entity"

func (uc *CreateUserUseCase) CreateUser(user *entity.User) error {
	err := uc.UserGateway.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}
