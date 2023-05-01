package documentationRepository

import (
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (r *DocumentationRepository) CreateDocumentation(payload documentationDTO.CreateDocumentationInputDTO, userID uint) error {
	documentation := entity.Documentation{
		Title:    payload.Title,
		Content:  payload.Content,
		AuthorID: userID,
	}

	if err := r.DB.Create(&documentation).Error; err != nil {
		return err
	}

	return nil
}
