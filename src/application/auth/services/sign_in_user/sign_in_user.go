package authService

import (
	"fmt"
	"strings"

	authDtos "github.com/rodrigoRVSN/beeus-api/src/application/auth/dtos"
	helpers "github.com/rodrigoRVSN/beeus-api/src/application/auth/services/sign_in_user/helpers"
	hash "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
	fieldsValidator "github.com/rodrigoRVSN/beeus-api/src/infra/helpers/validator"
)

func SignInUser(payload *authDtos.SignInInput) (*authDtos.SignInOutput, error) {
	errors := fieldsValidator.ValidateStruct(payload)

	if errors != nil {
		return nil, errors
	}

	user, err := helpers.FindUserByEmail(strings.ToLower(payload.Email))
	isPasswordCorrect := hash.CheckPasswordHash(payload.Password, user.Password)

	if err != nil || !isPasswordCorrect {
		return nil, fmt.Errorf("email ou senha incorretos")
	}

	tokenString, err := helpers.CreateJwtToken(user.Id)

	if err != nil {
		return nil, fmt.Errorf("erro ao gerar token JWT: %v", err)
	}

	return &authDtos.SignInOutput{
		Token: tokenString,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
