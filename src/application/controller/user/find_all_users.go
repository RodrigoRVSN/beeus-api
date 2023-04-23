package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (controller *UserController) FindAllUsers(context *fiber.Ctx) error {
	user := new(entity.User)

	users, err := controller.useCase.FindAllUsers(user)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	return context.Status(fiber.StatusOK).JSON(users)
}
