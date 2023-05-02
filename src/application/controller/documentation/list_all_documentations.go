package documentationController

import (
	"github.com/gofiber/fiber/v2"
)

func (controller *DocumentationController) ListAllDocumentations(context *fiber.Ctx) error {
	documentations, err := controller.useCase.ListAllDocumentations()

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	return context.Status(fiber.StatusOK).JSON(documentations)
}
