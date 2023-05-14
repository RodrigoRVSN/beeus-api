package documentationRepository

import (
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (r *DocumentationRepository) EditDocumentation(payload documentationDTO.EditDocumentationInputDTO, tags []entity.Tag, documentationID uint) (*documentationDTO.EditDocumentationOutputDTO, error) {
	documentation := entity.Documentation{
		Id:      documentationID,
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    tags,
	}

	if err := r.DB.Model(&documentation).Where("id = ?", documentationID).Updates(&documentation).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Preload("Author").First(&documentation, documentationID).Error; err != nil {
		return nil, err
	}

	return &documentationDTO.EditDocumentationOutputDTO{
		Id:        documentationID,
		Title:     documentation.Title,
		Content:   documentation.Content,
		CreatedAt: documentation.CreatedAt,
		UpdatedAt: documentation.UpdatedAt,
		Tags:      tags,
		Author: userDTO.UserWithoutPasswordDTO{
			Id:        documentation.Author.Id,
			Name:      documentation.Author.Name,
			Email:     documentation.Author.Email,
			AvatarUrl: documentation.Author.AvatarUrl,
			Points:    documentation.Author.Points,
			CreatedAt: documentation.Author.CreatedAt,
		},
	}, nil
}
