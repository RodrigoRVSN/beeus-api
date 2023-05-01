package documentationUseCase

import (
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
)

func (uc *DocumentationUseCase) CreateDocumentation(payload documentationDTO.CreateDocumentationInputDTO, userID uint) (*documentationDTO.CreateDocumentationOutputDTO, error) {
	user, userWasNotFound := uc.UserUseCase.GetUserById(userID)

	if userWasNotFound != nil {
		return nil, userWasNotFound
	}

	documentation, err := uc.DocumentationGateway.CreateDocumentation(payload, user)

	if err != nil {
		return nil, err
	}

	return documentation, err
}
