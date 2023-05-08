package documentationUseCase

func (uc *DocumentationUseCase) DeleteDocumentation(documentationID uint) error {
	documentation, err := uc.DocumentationGateway.FindDocumentationById(documentationID)

	if err != nil {
		return err
	}

	if err := uc.DocumentationGateway.DeleteDocumentation(documentation); err != nil {
		return err
	}

	return nil
}
