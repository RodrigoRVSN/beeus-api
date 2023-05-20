package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/rodrigoRVSN/beeus-api/src/infra/context"
)

type CustomClaims struct {
	jwt.StandardClaims
	UserId uint `json:"sub"`
}

func AuthMiddleware(ctx context.Context) error {
	token := ctx.GetHeader("Authorization")

	if token == "" {
		return ctx.SendJson(fiber.StatusUnauthorized, fiber.Map{
			"message": "Token de autenticação ausente",
		})
	}

	token = strings.ReplaceAll(token, "Bearer ", "")

	claims := &CustomClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return ctx.SendJson(fiber.StatusUnauthorized, fiber.Map{
			"message": "Token de autenticação inválido",
		})
	}

	ctx.SetMiddlewareParam("userId", claims.UserId)

	return ctx.Next()
}
