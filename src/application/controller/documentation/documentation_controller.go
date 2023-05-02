package documentationController

import (
	documentationUseCase "github.com/rodrigoRVSN/beeus-api/src/application/use_case/documentation"
)

type DocumentationController struct {
	useCase documentationUseCase.DocumentationUseCase
}

func NewDocumentationController(useCase documentationUseCase.DocumentationUseCase) *DocumentationController {
	return &DocumentationController{
		useCase: useCase,
	}
}
