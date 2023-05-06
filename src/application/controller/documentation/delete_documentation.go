package documentationController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (controller *DocumentationController) DeleteDocumentation(context *fiber.Ctx) error {
	documentationIDStr := context.Params("documentationId")
	documentationID, err := strconv.ParseUint(documentationIDStr, 10, 64)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).SendString("Invalid documentation ID")
	}

	if err := controller.useCase.DeleteDocumentation(uint(documentationID)); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return context.Status(fiber.StatusNoContent).SendString("Documentação deletada com sucesso!")
}
