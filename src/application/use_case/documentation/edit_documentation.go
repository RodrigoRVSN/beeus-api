package documentationUseCase

import (
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (uc *DocumentationUseCase) EditDocumentation(payload documentationDTO.EditDocumentationInputDTO, documentationID uint) (*documentationDTO.EditDocumentationOutputDTO, error) {
	tags := make([]entity.Tag, len(payload.Tags))
	for i, actualTag := range payload.Tags {
		tag, err := uc.TagRepository.FindTagByName(actualTag)

		if err != nil {
			return nil, err
		}

		if tag != nil {
			tags[i] = entity.Tag{
				Id:   tag.Id,
				Name: tag.Name,
			}
			continue
		}

		createdTag, createdError := uc.TagRepository.CreateTag(actualTag)

		if createdError != nil {
			return nil, err
		}

		tags[i] = entity.Tag{
			Id:   createdTag.Id,
			Name: createdTag.Name,
		}
	}

	documentation, err := uc.DocumentationGateway.EditDocumentation(payload, tags, documentationID)

	if err != nil {
		return nil, err
	}

	return documentation, err
}