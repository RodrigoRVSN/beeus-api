package gateway

import (
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
)

type DocumentationGateway interface {
	CreateDocumentation(payload documentationDTO.CreateDocumentationInputDTO, userID uint) error
	ListAllDocumentations() ([]documentationDTO.ListDocumentationsOutputDTO, error)
}
