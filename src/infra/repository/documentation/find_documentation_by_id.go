package documentationRepository

import (
	"fmt"

	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (r *DocumentationRepository) FindDocumentationById(documentationId uint) (*documentationDTO.FindDocumentationByIdOutputDTO, error) {
	documentation := entity.Documentation{}

	if err := r.DB.Preload("Author").Find(&documentation).Error; err != nil {
		if err.Error() == "record not found" {
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
