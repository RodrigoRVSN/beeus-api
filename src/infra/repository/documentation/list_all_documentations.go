package documentationRepository

import (
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (r *DocumentationRepository) ListAllDocumentations() ([]documentationDTO.ListDocumentationsOutputDTO, error) {
	documentations := []entity.Documentation{}

	if err := r.DB.Preload("Tags").Preload("Author").Find(&documentations).Error; err != nil {
		return nil, err
	}

	outputDTOs := make([]documentationDTO.ListDocumentationsOutputDTO, len(documentations))
	for i, doc := range documentations {
		outputDTOs[i] = documentationDTO.ListDocumentationsOutputDTO{
			Id:        doc.Id,
			Title:     doc.Title,
			Content:   doc.Content,
			CreatedAt: doc.CreatedAt,
			UpdatedAt: doc.UpdatedAt,
			Tags:      doc.Tags,
			Author:    userDTO.MapUserWithoutPasswordToDTO(doc.Author),
		}
	}

	return outputDTOs, nil
}
