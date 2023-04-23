package userRepository

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rodrigoRVSN/beeus-api/src/application/dto"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
	hash "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
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

func (r *UserRepository) SignInUser(payload dto.SignInInputDTO) (*dto.SignInOutputDTO, error) {
	user := entity.User{}

	email := strings.ToLower(payload.Email)
	userFound := r.DB.First(&user, "email = ?", email).Error

	isPasswordCorrect := hash.CheckPasswordHash(payload.Password, user.Password)

	if userFound != nil || !isPasswordCorrect {
		return nil, fmt.Errorf("email ou senha incorretos")
	}

	tokenString, err := createJwtToken(user.Id)

	if err != nil {
		return nil, fmt.Errorf("erro ao gerar token")
	}

	return &dto.SignInOutputDTO{
		Token: tokenString,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
