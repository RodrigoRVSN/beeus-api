package routes

import (
	"github.com/gofiber/fiber/v2"
	documentationController "github.com/rodrigoRVSN/beeus-api/src/application/controller/documentation"
	"github.com/rodrigoRVSN/beeus-api/src/infra/context"
	"github.com/rodrigoRVSN/beeus-api/src/infra/middleware"
)

func DocumentationRoutes(app *fiber.App, documentationController documentationController.DocumentationControllerInterface) {
	docRoutes := app.Group("/documentation")

	docRoutes.Get("", context.AdaptHandler(documentationController.ListAllDocumentations))

	docRoutes.Get("/:documentationId", context.AdaptHandler(middleware.AuthMiddleware), context.AdaptHandler(documentationController.GetDocumentation))

	docRoutes.Post("", context.AdaptHandler(middleware.AuthMiddleware), context.AdaptHandler(documentationController.CreateDocumentation))

	docRoutes.Delete("/:documentationId", context.AdaptHandler(middleware.AuthMiddleware), context.AdaptHandler(documentationController.DeleteDocumentation))

	docRoutes.Put("/:documentationId", context.AdaptHandler(middleware.AuthMiddleware), context.AdaptHandler(documentationController.EditDocumentation))
}
