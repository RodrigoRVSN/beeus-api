package userUseCase

import "github.com/rodrigoRVSN/beeus-api/src/domain/entity"

func (uc *CreateUserUseCase) CreateUser(user *entity.User) error {
	userWasFound := uc.UserGateway.GetUserByEmail(user, user.Email)

	if userWasFound != nil {
		return userWasFound
	}

	err := uc.UserGateway.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}
