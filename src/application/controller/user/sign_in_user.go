package userController

import (
	"github.com/gofiber/fiber/v2"
	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
	"github.com/rodrigoRVSN/beeus-api/src/infra/context"
	fieldsValidator "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
)

func (controller *UserController) SignInUser(ctx context.Context) error {
	payload := new(userDTO.SignInInputDTO)

	if err := ctx.ParseBody(&payload); err != nil {
		return ctx.SendJson(fiber.StatusBadRequest, fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := fieldsValidator.ValidateStruct(*payload)

	if errors != nil {
		return ctx.SendJson(fiber.StatusBadRequest, errors)
	}

	user, err := controller.useCase.SignInUser(*payload)

	if err != nil {
		return ctx.SendJson(fiber.StatusUnauthorized, fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.SendJson(fiber.StatusOK, user)
}
