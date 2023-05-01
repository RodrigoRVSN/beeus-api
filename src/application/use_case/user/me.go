package userUseCase

import userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"

func (uc *UserUseCase) Me(userId uint) (*userDTO.FindUserByIdOutputDTO, error) {
	user, err := uc.UserGateway.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}
