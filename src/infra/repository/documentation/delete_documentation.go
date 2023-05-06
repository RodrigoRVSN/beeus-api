package documentationRepository

import (
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
)

func (r *DocumentationRepository) DeleteDocumentation(documentation *documentationDTO.FindDocumentationByIdOutputDTO) error {
	err := r.DB.Delete(documentation).Error

	return err
}
