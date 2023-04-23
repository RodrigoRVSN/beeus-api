package authController

import (
	"github.com/gofiber/fiber/v2"
	authService "github.com/rodrigoRVSN/beeus-api/src/application/auth/services"
	"github.com/rodrigoRVSN/beeus-api/src/application/users/entities"
)

func CreateUser(context *fiber.Ctx) error {
	payload := new(entities.User)

	if err := context.BodyParser(&payload); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	if err := authService.CreateUser(payload); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	return context.Status(fiber.StatusCreated).SendString("User registered!")
}
