package documentationRepository

import (
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (r *DocumentationRepository) CreateDocumentation(payload documentationDTO.CreateDocumentationInputDTO, user *userDTO.FindUserByIdOutputDTO) (*documentationDTO.CreateDocumentationOutputDTO, error) {
	documentation := entity.Documentation{
		Title:    payload.Title,
		Content:  payload.Content,
		AuthorID: user.Id,
	}

	if err := r.DB.Create(&documentation).Error; err != nil {
		return nil, err
	}

	return &documentationDTO.CreateDocumentationOutputDTO{
		Id:        documentation.Id,
		Title:     documentation.Title,
		Content:   documentation.Content,
		CreatedAt: documentation.CreatedAt,
		UpdatedAt: documentation.UpdatedAt,
		Author: userDTO.UserWithoutPasswordDTO{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			AvatarUrl: user.AvatarUrl,
			Points:    user.Points,
			CreatedAt: user.CreatedAt,
		},
	}, nil
}
