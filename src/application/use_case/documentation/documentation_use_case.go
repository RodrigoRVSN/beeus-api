package documentationUseCase

import (
	userUseCase "github.com/rodrigoRVSN/beeus-api/src/application/use_case/user"
	"github.com/rodrigoRVSN/beeus-api/src/domain/gateway"
)

type DocumentationUseCase struct {
	DocumentationGateway gateway.DocumentationGateway
	UserUseCase          gateway.UserGateway
	TagRepository        gateway.TagGateway
}

func NewDocumentationUseCase(documentationGateway gateway.DocumentationGateway, userUseCase *userUseCase.UserUseCase, tagRepository gateway.TagGateway) *DocumentationUseCase {
	return &DocumentationUseCase{
		DocumentationGateway: documentationGateway,
		UserUseCase:          userUseCase.UserGateway,
		TagRepository:        tagRepository,
	}
}
