package gateway

import (
	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

type UserGateway interface {
	GetUserByEmail(email string) error
	CreateUser(payload userDTO.SignUpInputDTO) error
	SignInUser(payload userDTO.SignInInputDTO) (*userDTO.SignInOutputDTO, error)
	GetUserById(userId uint) (*userDTO.FindUserByIdOutputDTO, error)
	// TODO: The FindAllUsers method is just a health check and will be removed soon. The architecture is not following the patterns because it doesnt matters for now.
	FindAllUsers(user *entity.User) ([]entity.User, error)
}
