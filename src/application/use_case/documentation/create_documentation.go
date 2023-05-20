package documentationUseCase

import (
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (uc *DocumentationUseCase) CreateDocumentation(payload documentationDTO.CreateDocumentationInputDTO, userID uint) (*documentationDTO.CreateDocumentationOutputDTO, error) {
	user, userWasNotFound := uc.UserUseCase.GetUserById(userID)

	if userWasNotFound != nil {
		return nil, userWasNotFound
	}

	tags := make([]entity.Tag, len(payload.Tags))
	for i, actualTag := range payload.Tags {
		tag, err := uc.TagRepository.FindTagByName(actualTag)

		if err != nil {
			return nil, err
		}

		if tag != nil {
			tags[i] = *entity.NewTag(tag.Id, tag.Name)
			continue
		}

		createdTag, createdError := uc.TagRepository.CreateTag(actualTag)

		if createdError != nil {
			return nil, err
		}

		tags[i] = *entity.NewTag(createdTag.Id, createdTag.Name)
	}

	documentation, err := uc.DocumentationGateway.CreateDocumentation(payload, user, tags)

	if err != nil {
		return nil, err
	}

	return documentation, err
}
