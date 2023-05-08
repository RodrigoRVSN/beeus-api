package documentationUseCase

import (
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
)

func (uc *DocumentationUseCase) GetDocumentation(documentationID uint) (*documentationDTO.FindDocumentationByIdOutputDTO, error) {
	documentation, err := uc.DocumentationGateway.FindDocumentationById(documentationID)

	if err != nil {
		return nil, err
	}

	return documentation, nil
}
