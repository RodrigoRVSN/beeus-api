package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	jwt.StandardClaims
	UserId uint `json:"sub"`
}

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token de autenticação ausente",
		})
	}

	token = strings.ReplaceAll(token, "Bearer ", "")

	claims := &CustomClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token de autenticação inválido",
		})
	}

	c.Locals("userId", claims.UserId)

	return c.Next()
}
