package documentationController

import (
	"github.com/gofiber/fiber/v2"
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
	"github.com/rodrigoRVSN/beeus-api/src/infra/context"
	fieldsValidator "github.com/rodrigoRVSN/beeus-api/src/infra/helpers"
)

func (controller *DocumentationController) CreateDocumentation(ctx context.Context) error {
	userID := ctx.GetMiddlewareParam("userId")
	test := userID.(uint)
	payload := new(documentationDTO.CreateDocumentationInputDTO)

	if err := ctx.ParseBody(&payload); err != nil {
		return ctx.SendJson(fiber.StatusBadRequest, fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := fieldsValidator.ValidateStruct(*payload)

	if errors != nil {
		return ctx.SendJson(fiber.StatusBadRequest, errors)
	}

	documentation, err := controller.useCase.CreateDocumentation(*payload, test)

	if err != nil {
		return ctx.SendJson(fiber.StatusBadRequest, fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.SendJson(fiber.StatusCreated, documentation)
}
