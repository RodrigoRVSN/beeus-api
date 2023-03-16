package userController

import (
	"github.com/gofiber/fiber/v2"
	database "github.com/rodrigoRVSN/beeus-api/src/infra/config"
	"github.com/rodrigoRVSN/beeus-api/src/models/users/entities"
)

func ListUsers(context *fiber.Ctx) error {
	users := []entities.User{}

	database.DB.Db.Find(&users)

	return context.Status(200).JSON(users)
}

func CreateUser(context *fiber.Ctx) error {
	user := new(entities.User)

	if err := context.BodyParser(user); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&user)

	return context.Status(200).JSON(user)
}
