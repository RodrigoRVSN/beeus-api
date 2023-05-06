package documentationRepository

import (
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (r *DocumentationRepository) DeleteDocumentation(documentation *documentationDTO.FindDocumentationByIdOutputDTO) error {
	err := r.DB.Delete(&entity.Documentation{}, documentation.Id).Error

	return err
}
