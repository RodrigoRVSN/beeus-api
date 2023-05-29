package gateway

import (
	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
)

type UserGateway interface {
	GetUserByEmail(email string) error
	CreateUser(payload userDTO.SignUpInputDTO) error
	SignInUser(payload userDTO.SignInInputDTO) (*userDTO.SignInOutputDTO, error)
	GetUserById(userId uint) (*userDTO.FindUserByIdOutputDTO, error)
	FindAllUsers(sortBy string, sortOrder string) (*[]userDTO.UserWithoutPasswordDTO, error)
}
