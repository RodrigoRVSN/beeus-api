package documentationController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
	fieldsValidator "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
)

func (controller *DocumentationController) EditDocumentation(ctx *fiber.Ctx) error {
	documentationIDStr := ctx.Params("documentationId")
	documentationID, err := strconv.ParseUint(documentationIDStr, 10, 64)
	payload := new(documentationDTO.EditDocumentationInputDTO)

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := fieldsValidator.ValidateStruct(*payload)

	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	documentation, err := controller.useCase.EditDocumentation(*payload, uint(documentationID))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(documentation)
}
