package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoRVSN/beeus-api/src/models/users/entities"
	userService "github.com/rodrigoRVSN/beeus-api/src/models/users/services"
)

func CreateUser(context *fiber.Ctx) error {
	payload := new(entities.User)

	if err := context.BodyParser(&payload); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	if err := userService.CreateUser(payload); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	return context.Status(fiber.StatusCreated).SendString("User registered!")
}
