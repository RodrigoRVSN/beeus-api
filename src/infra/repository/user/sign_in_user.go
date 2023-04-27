package userRepository

import (
	"fmt"
	"strings"

	"github.com/rodrigoRVSN/beeus-api/src/application/dto"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
	hash "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
	"github.com/rodrigoRVSN/beeus-api/src/infra/service"
)

func (r *UserRepository) SignInUser(payload dto.SignInInputDTO) (*dto.SignInOutputDTO, error) {
	user := entity.User{}

	email := strings.ToLower(payload.Email)
	userFound := r.DB.First(&user, "email = ?", email).Error

	isPasswordCorrect := hash.CheckPasswordHash(payload.Password, user.Password)

	if userFound != nil || !isPasswordCorrect {
		return nil, fmt.Errorf("email ou senha incorretos")
	}

	tokenString, err := service.CreateJwtToken(user.Id)

	if err != nil {
		return nil, fmt.Errorf("erro ao gerar token")
	}

	return &dto.SignInOutputDTO{
		Token: tokenString,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
