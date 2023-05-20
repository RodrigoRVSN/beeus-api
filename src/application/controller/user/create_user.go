package userController

import (
	"github.com/gofiber/fiber/v2"
	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
	"github.com/rodrigoRVSN/beeus-api/src/infra/context"
	fieldsValidator "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
)

func (controller *UserController) CreateUser(ctx context.Context) error {
	payload := new(userDTO.SignUpInputDTO)

	if err := ctx.ParseBody(&payload); err != nil {
		return ctx.SendJson(fiber.StatusBadRequest, fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := fieldsValidator.ValidateStruct(*payload)

	if errors != nil {
		return ctx.SendJson(fiber.StatusBadRequest, errors)
	}

	err := controller.useCase.CreateUser(*payload)

	if err != nil {
		return ctx.SendJson(fiber.StatusBadRequest, fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.SendJson(fiber.StatusCreated, "User registered!")
}
