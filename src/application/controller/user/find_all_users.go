package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
	"github.com/rodrigoRVSN/beeus-api/src/infra/context"
)

func (controller *UserController) FindAllUsers(ctx context.Context) error {
	user := new(entity.User)

	users, err := controller.useCase.FindAllUsers(user)

	if err != nil {
		return ctx.SendJson(fiber.StatusBadRequest, err)
	}

	return ctx.SendJson(fiber.StatusOK, users)
}
