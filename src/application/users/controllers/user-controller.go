package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoRVSN/beeus-api/src/application/users/entities"
	database "github.com/rodrigoRVSN/beeus-api/src/infra/config"
)

func ListUsers(context *fiber.Ctx) error {
	users := []entities.User{}

	database.DB.Find(&users)

	return context.Status(200).JSON(users)
}
