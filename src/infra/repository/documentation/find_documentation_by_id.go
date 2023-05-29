package documentationRepository

import (
	"fmt"

	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
	"gorm.io/gorm"
)

func (r *DocumentationRepository) FindDocumentationById(documentationId uint) (*documentationDTO.FindDocumentationByIdOutputDTO, error) {
	documentation := entity.Documentation{}

	if err := r.DB.Preload("Tags").Preload("Author").First(&documentation, documentationId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("documentação não encontrada")
		}
		return nil, err
	}

	return &documentationDTO.FindDocumentationByIdOutputDTO{
		Id:        documentation.Id,
		Title:     documentation.Title,
		Content:   documentation.Content,
		CreatedAt: documentation.CreatedAt,
		UpdatedAt: documentation.UpdatedAt,
		Tags:      documentation.Tags,
		Author:    userDTO.MapUserWithoutPasswordToDTO(documentation.Author),
	}, nil
}
