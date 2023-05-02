package userUseCase

import "github.com/rodrigoRVSN/beeus-api/src/domain/entity"

func (uc *UserUseCase) FindAllUsers(user *entity.User) ([]entity.User, error) {
	users, err := uc.UserGateway.FindAllUsers(user)

	if err != nil {
		return nil, err
	}

	return users, nil
}
