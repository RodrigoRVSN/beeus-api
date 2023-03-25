package authService

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	authDtos "github.com/rodrigoRVSN/beeus-api/src/app/auth/dtos"
	"github.com/rodrigoRVSN/beeus-api/src/app/users/entities"
	database "github.com/rodrigoRVSN/beeus-api/src/infra/config"
	hash "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
	fieldsValidator "github.com/rodrigoRVSN/beeus-api/src/infra/helpers/validator"
)

func SignInUser(payload *authDtos.SignInInput) (*authDtos.SignInOutput, error) {
	errors := fieldsValidator.ValidateStruct(payload)

	if errors != nil {
		return nil, errors
	}

	var user entities.User
	result := database.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	isPasswordCorrect := hash.CheckPasswordHash(payload.Password, user.Password)

	if result.Error != nil || !isPasswordCorrect {
		return nil, fmt.Errorf("email ou senha incorretos")
	}

	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)

	expiredIn, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))

	claims["sub"] = user.ID
	claims["exp"] = now.Add(time.Duration(expiredIn)).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return nil, fmt.Errorf("erro ao gerar token JWT: %v", err)
	}

	return &authDtos.SignInOutput{
		Token: tokenString,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
