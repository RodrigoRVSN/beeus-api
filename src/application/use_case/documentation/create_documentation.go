package documentationUseCase

import (
	"github.com/rodrigoRVSN/beeus-api/src/application/dto"
)

func (uc *DocumentationUseCase) CreateDocumentation(payload dto.CreateDocumentationInputDTO, userID uint) error {
	user, userWasNotFound := uc.UserUseCase.GetUserById(userID)

	if userWasNotFound != nil {
		return userWasNotFound
	}

	err := uc.DocumentationGateway.CreateDocumentation(payload, user.Id)

	if err != nil {
		return err
	}

	return nil
}
