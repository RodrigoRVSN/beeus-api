package documentationController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (controller *DocumentationController) GetDocumentation(context *fiber.Ctx) error {
	documentationIDStr := context.Params("documentationId")
	documentationID, err := strconv.ParseUint(documentationIDStr, 10, 64)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).SendString("Invalid documentation ID")
	}

	documentation, err := controller.useCase.GetDocumentation(uint(documentationID))
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return context.Status(fiber.StatusOK).JSON(documentation)
}
