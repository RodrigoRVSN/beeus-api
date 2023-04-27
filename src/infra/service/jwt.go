package service

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJwtToken(userID uint) (string, error) {
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
