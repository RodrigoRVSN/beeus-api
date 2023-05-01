package documentationController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
	fieldsValidator "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
)

func (controller *DocumentationController) CreateDocumentation(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(uint)
	payload := new(documentationDTO.CreateDocumentationInputDTO)

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := fieldsValidator.ValidateStruct(*payload)

	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	err := controller.useCase.CreateDocumentation(*payload, userID)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).SendString("Documentation registered!")
}
