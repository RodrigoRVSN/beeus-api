package documentationController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoRVSN/beeus-api/src/infra/context"
)

func (controller *DocumentationController) ListAllDocumentations(ctx context.Context) error {
	documentations, err := controller.useCase.ListAllDocumentations()

	if err != nil {
		return ctx.SendJson(fiber.StatusBadRequest, err)
	}

	return ctx.SendJson(fiber.StatusOK, documentations)
}
