package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
	fieldsValidator "github.com/rodrigoRVSN/beeus-api/src/infra/helpers/validator"
)

func (controller *UserController) CreateUser(context *fiber.Ctx) error {
	payload := new(entity.User)

	if err := context.BodyParser(&payload); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := fieldsValidator.ValidateStruct(*payload)

	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	err := controller.useCase.CreateUser(payload)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return context.Status(fiber.StatusCreated).SendString("User registered!")
}
