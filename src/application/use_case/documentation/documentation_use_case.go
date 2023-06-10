package documentationUseCase

import (
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
	"github.com/rodrigoRVSN/beeus-api/src/domain/gateway"
)

type DocumentationUseCaseInterface interface {
	CreateDocumentation(payload documentationDTO.CreateDocumentationInputDTO, userID uint) (*documentationDTO.CreateDocumentationOutputDTO, error)
	ListAllDocumentations() ([]documentationDTO.ListDocumentationsOutputDTO, error)
	GetDocumentation(documentationID uint) (*documentationDTO.FindDocumentationByIdOutputDTO, error)
	EditDocumentation(payload documentationDTO.EditDocumentationInputDTO, documentationID uint) (*documentationDTO.EditDocumentationOutputDTO, error)
	DeleteDocumentation(documentationID uint) error
}

type DocumentationUseCase struct {
	DocumentationGateway gateway.DocumentationGateway
	UserGateway          gateway.UserGateway
	TagGateway           gateway.TagGateway
}

func NewDocumentationUseCase(documentationGateway gateway.DocumentationGateway, userGateway gateway.UserGateway, tagGateway gateway.TagGateway) DocumentationUseCaseInterface {
	return &DocumentationUseCase{
		DocumentationGateway: documentationGateway,
		UserGateway:          userGateway,
		TagGateway:           tagGateway,
	}
}
