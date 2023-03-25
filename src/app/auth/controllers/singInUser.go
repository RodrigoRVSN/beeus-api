package authController

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	userDtos "github.com/rodrigoRVSN/beeus-api/src/app/auth/dtos"
	userService "github.com/rodrigoRVSN/beeus-api/src/app/auth/services"
)

func SignInUser(context *fiber.Ctx) error {
	payload := new(userDtos.SignInInput)

	if err := context.BodyParser(&payload); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	user, err := userService.SignInUser(payload)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	maxAge, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))

	context.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    user.Token,
		Path:     "/",
		MaxAge:   maxAge,
		Secure:   false,
		HTTPOnly: true,
	})

	return context.Status(fiber.StatusOK).JSON(user)
}
