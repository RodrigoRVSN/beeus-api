package gateway

import (
	"github.com/rodrigoRVSN/beeus-api/src/application/dto"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

type UserGateway interface {
	CheckIfUserAlreadyExists(payload dto.SignUpInputDTO, email string) error
	CreateUser(payload dto.SignUpInputDTO) error
	SignInUser(payload dto.SignInInputDTO) (*dto.SignInOutputDTO, error)
	GetUserById(userId uint) (*dto.FindUserByIdOutputDTO, error)
	// TODO: The FindAllUsers method is just a health check and will be removed soon. The architecture is not following the patterns because it doesnt matters for now.
	FindAllUsers(user *entity.User) ([]entity.User, error)
}
