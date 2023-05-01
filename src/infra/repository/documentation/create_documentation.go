package documentationRepository

import (
	"github.com/rodrigoRVSN/beeus-api/src/application/dto"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (r *DocumentationRepository) CreateDocumentation(payload dto.CreateDocumentationInputDTO, userID uint) error {
	documentation := entity.Documentation{
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  userID,
	}

	if err := r.DB.Create(&documentation).Error; err != nil {
		return err
	}

	return nil
}
