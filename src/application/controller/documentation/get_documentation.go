package documentationController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoRVSN/beeus-api/src/infra/context"
)

func (controller *DocumentationController) GetDocumentation(ctx context.Context) error {
	documentationIDStr := ctx.GetParam("documentationId")
	documentationID, err := strconv.ParseUint(documentationIDStr, 10, 64)

	if err != nil {
		return ctx.SendJson(fiber.StatusBadRequest, fiber.Map{"status": "fail", "message": "ID de documentação inválido"})
	}

	documentation, err := controller.useCase.GetDocumentation(uint(documentationID))
	if err != nil {
		return ctx.SendJson(fiber.StatusBadRequest, fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.SendJson(fiber.StatusOK, documentation)
}
