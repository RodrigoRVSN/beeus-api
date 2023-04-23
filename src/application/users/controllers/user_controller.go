package userController

import (
	"github.com/gofiber/fiber/v2"
	database "github.com/rodrigoRVSN/beeus-api/src/config"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func ListUsers(context *fiber.Ctx) error {
	users := []entity.User{}

	database.DB.Find(&users)

	return context.Status(200).JSON(users)
}
