package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoRVSN/beeus-api/src/infra/context"
)

func (controller *UserController) GetRankedUsers(ctx context.Context) error {
	users, err := controller.useCase.GetRankedUsers()

	if err != nil {
		return ctx.SendJson(fiber.StatusBadRequest, fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.SendJson(fiber.StatusOK, users)
}
