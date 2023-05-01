package documentationRepository

import (
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (r *DocumentationRepository) ListAllDocumentations() ([]entity.Documentation, error) {
	documentations := []entity.Documentation{}

	response := r.DB.Find(&documentations)

	if response.Error != nil {
		return nil, response.Error
	}

	return documentations, nil
}
