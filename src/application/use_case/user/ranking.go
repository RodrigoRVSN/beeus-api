package userUseCase

import userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"

func (uc *UserUseCase) GetRankedUsers() (*[]userDTO.UserWithoutPasswordDTO, error) {
	users, err := uc.UserGateway.FindAllUsers("points", "desc")

	if err != nil {
		return nil, err
	}

	return users, nil
}
