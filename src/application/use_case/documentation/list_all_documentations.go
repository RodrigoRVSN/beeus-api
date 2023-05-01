package documentationUseCase

import (
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (uc *DocumentationUseCase) ListAllDocumentations() ([]entity.Documentation, error) {
	documentations, err := uc.DocumentationGateway.ListAllDocumentations()

	if err != nil {
		return nil, err
	}

	return documentations, err
}
