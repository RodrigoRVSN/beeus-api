package authController

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	authDtos "github.com/rodrigoRVSN/beeus-api/src/app/auth/dtos"
	authService "github.com/rodrigoRVSN/beeus-api/src/app/auth/services"
)

func injectTokenCookie(context *fiber.Ctx, user *authDtos.SignInOutput) {
	maxAge, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))

	context.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    user.Token,
		Path:     "/",
		MaxAge:   maxAge,
		Secure:   false,
		HTTPOnly: true,
	})
}

func SignInUser(context *fiber.Ctx) error {
	payload := new(authDtos.SignInInput)

	if err := context.BodyParser(&payload); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	user, err := authService.SignInUser(payload)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	injectTokenCookie(context, user)

	return context.Status(fiber.StatusOK).JSON(user)
}
