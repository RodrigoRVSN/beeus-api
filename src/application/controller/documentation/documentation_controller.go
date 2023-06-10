package documentationController

import (
	documentationUseCase "github.com/rodrigoRVSN/beeus-api/src/application/use_case/documentation"
	"github.com/rodrigoRVSN/beeus-api/src/infra/context"
)

type DocumentationControllerInterface interface {
	CreateDocumentation(ctx context.Context) error
	ListAllDocumentations(ctx context.Context) error
	GetDocumentation(ctx context.Context) error
	EditDocumentation(ctx context.Context) error
	DeleteDocumentation(ctx context.Context) error
}

type DocumentationController struct {
	useCase documentationUseCase.DocumentationUseCaseInterface
}

func NewDocumentationController(useCase documentationUseCase.DocumentationUseCaseInterface) DocumentationControllerInterface {
	return &DocumentationController{
		useCase: useCase,
	}
}
