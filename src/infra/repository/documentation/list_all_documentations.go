package documentationRepository

import (
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (r *DocumentationRepository) ListAllDocumentations() ([]entity.Documentation, error) {
	documentations := []entity.Documentation{}

	if err := r.DB.Preload("User").Find(&documentations).Error; err != nil {
		return nil, err
	}

	return documentations, nil
}
