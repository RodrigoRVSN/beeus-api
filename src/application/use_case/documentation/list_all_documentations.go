package documentationUseCase

import documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"

func (uc *DocumentationUseCase) ListAllDocumentations() ([]documentationDTO.ListDocumentationsOutputDTO, error) {
	documentations, err := uc.DocumentationGateway.ListAllDocumentations()

	if err != nil {
		return nil, err
	}

	return documentations, err
}
