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
	useCase documentationUseCase.UserControllerInterface
}

func NewDocumentationController(useCase documentationUseCase.UserControllerInterface) DocumentationControllerInterface {
	return &DocumentationController{
		useCase: useCase,
	}
}
