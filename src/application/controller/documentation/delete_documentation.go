package documentationController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoRVSN/beeus-api/src/infra/context"
)

func (controller *DocumentationController) DeleteDocumentation(ctx context.Context) error {
	documentationIDStr := ctx.GetParam("documentationId")
	documentationID, err := strconv.ParseUint(documentationIDStr, 10, 64)

	if err != nil {
		return ctx.SendJson(fiber.StatusBadRequest, fiber.Map{"status": "fail", "message": "ID de documentação inválido"})
	}

	if err := controller.useCase.DeleteDocumentation(uint(documentationID)); err != nil {
		return ctx.SendJson(fiber.StatusBadRequest, fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.SendJson(fiber.StatusNoContent, "Documentação deletada com sucesso!")
}
