package userService

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	database "github.com/rodrigoRVSN/beeus-api/src/infra/config"
	hash "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
	fieldsValidator "github.com/rodrigoRVSN/beeus-api/src/infra/helpers/validator"
	userDtos "github.com/rodrigoRVSN/beeus-api/src/models/users/dtos"
	"github.com/rodrigoRVSN/beeus-api/src/models/users/entities"
)

func SignInUser(payload *userDtos.SignInInput) (*userDtos.SignInOutput, error) {
	errors := fieldsValidator.ValidateStruct(payload)

	if errors != nil {
		return nil, errors
	}

	var user entities.User
	result := database.DB.Db.First(&user, "email = ?", strings.ToLower(payload.Email))
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

	return &userDtos.SignInOutput{
		Token: tokenString,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
