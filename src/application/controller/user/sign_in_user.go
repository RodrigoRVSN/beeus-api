package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoRVSN/beeus-api/src/application/dto"
	fieldsValidator "github.com/rodrigoRVSN/beeus-api/src/infra/helpers/validator"
)

func (controller *UserController) SignInUser(context *fiber.Ctx) error {
	payload := new(dto.SignInInputDTO)

	if err := context.BodyParser(&payload); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := fieldsValidator.ValidateStruct(*payload)

	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	user, err := controller.useCase.SignInUser(*payload)

	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return context.Status(fiber.StatusOK).JSON(user)
}
