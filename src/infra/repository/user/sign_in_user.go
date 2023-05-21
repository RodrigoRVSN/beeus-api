package userRepository

import (
	"fmt"
	"strings"

	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
	hash "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
	"github.com/rodrigoRVSN/beeus-api/src/infra/service"
)

func (r *UserRepository) SignInUser(payload userDTO.SignInInputDTO) (*userDTO.SignInOutputDTO, error) {
	user := entity.User{}

	email := strings.ToLower(payload.Email)
	err := r.DB.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, fmt.Errorf("email ou senha incorretos")
	}

	isPasswordCorrect := hash.CheckPasswordHash(payload.Password, user.Password)
	if !isPasswordCorrect {
		return nil, fmt.Errorf("email ou senha incorretos")
	}

	tokenString, err := service.CreateJwtToken(user.Id)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar token")
	}

	return &userDTO.SignInOutputDTO{
		Token: tokenString,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
