package documentationController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
	"github.com/rodrigoRVSN/beeus-api/src/infra/context"
	fieldsValidator "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
)

func (controller *DocumentationController) EditDocumentation(ctx context.Context) error {
	documentationIDStr := ctx.GetParam("documentationId")
	documentationID, err := strconv.ParseUint(documentationIDStr, 10, 64)
	payload := new(documentationDTO.EditDocumentationInputDTO)

	if err != nil {
		return ctx.SendJson(fiber.StatusBadRequest, fiber.Map{"status": "fail", "message": "ID de documentação inválido"})
	}

	if err := ctx.ParseBody(&payload); err != nil {
		return ctx.SendJson(fiber.StatusBadRequest, fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := fieldsValidator.ValidateStruct(*payload)

	if errors != nil {
		return ctx.SendJson(fiber.StatusBadRequest, errors)
	}

	documentation, err := controller.useCase.EditDocumentation(*payload, uint(documentationID))

	if err != nil {
		return ctx.SendJson(fiber.StatusBadRequest, fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.SendJson(fiber.StatusOK, documentation)
}
