package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoRVSN/beeus-api/src/infra/context"
)

func (controller *UserController) Me(ctx context.Context) error {
	userId := ctx.GetMiddlewareParam("userId").(uint)

	user, err := controller.useCase.Me(userId)

	if err != nil {
		return ctx.SendJson(fiber.StatusBadRequest, fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.SendJson(fiber.StatusOK, user)
}
