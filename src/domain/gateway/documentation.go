package gateway

import (
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
)

type DocumentationGateway interface {
	CreateDocumentation(payload documentationDTO.CreateDocumentationInputDTO, user *userDTO.FindUserByIdOutputDTO) (*documentationDTO.CreateDocumentationOutputDTO, error)
	ListAllDocumentations() ([]documentationDTO.ListDocumentationsOutputDTO, error)
}
