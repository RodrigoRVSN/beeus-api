package userController

import (
	"github.com/gofiber/fiber/v2"
)

func (controller *UserController) Me(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)

	user, err := controller.useCase.Me(userId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}
