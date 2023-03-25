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

func createJwtToken(userID uint) (string, error) {
	now := time.Now().UTC()
	expiredIn, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))
	expirationTime := now.Add(time.Duration(expiredIn))

	claims := jwt.MapClaims{
		"sub": userID,
		"exp": expirationTime.Unix(),
		"iat": now.Unix(),
		"nbf": now.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func findUserByEmail(email string) (*entities.User, error) {
	var user entities.User

	if err := database.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func SignInUser(payload *authDtos.SignInInput) (*authDtos.SignInOutput, error) {
	errors := fieldsValidator.ValidateStruct(payload)

	if errors != nil {
		return nil, errors
	}

	user, err := findUserByEmail(strings.ToLower(payload.Email))
	isPasswordCorrect := hash.CheckPasswordHash(payload.Password, user.Password)

	if err != nil || !isPasswordCorrect {
		return nil, fmt.Errorf("email ou senha incorretos")
	}

	tokenString, err := createJwtToken(user.ID)

	if err != nil {
		return nil, fmt.Errorf("erro ao gerar token JWT: %v", err)
	}

	return &authDtos.SignInOutput{
		Token: tokenString,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
