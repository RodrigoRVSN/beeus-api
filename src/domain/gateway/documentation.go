package gateway

import (
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

type DocumentationGateway interface {
	CreateDocumentation(payload documentationDTO.CreateDocumentationInputDTO, user *userDTO.FindUserByIdOutputDTO, tags []entity.Tag) (*documentationDTO.CreateDocumentationOutputDTO, error)
	ListAllDocumentations() ([]documentationDTO.ListDocumentationsOutputDTO, error)
	DeleteDocumentation(documentation *documentationDTO.FindDocumentationByIdOutputDTO) error
	FindDocumentationById(documentationId uint) (*documentationDTO.FindDocumentationByIdOutputDTO, error)
	EditDocumentation(payload documentationDTO.EditDocumentationInputDTO, tags []entity.Tag, documentationID uint) (*documentationDTO.EditDocumentationOutputDTO, error)
}
