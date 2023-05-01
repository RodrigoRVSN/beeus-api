package gateway

import (
	"github.com/rodrigoRVSN/beeus-api/src/application/dto"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

type DocumentationGateway interface {
	CreateDocumentation(payload dto.CreateDocumentationInputDTO, userID uint) error
	ListAllDocumentations() ([]entity.Documentation, error)
}
