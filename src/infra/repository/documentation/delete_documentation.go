package documentationRepository

import (
	documentationDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/documentation"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (r *DocumentationRepository) DeleteDocumentation(documentation *documentationDTO.FindDocumentationByIdOutputDTO) error {
	doc := entity.Documentation{Id: documentation.Id}

	if err := r.DB.Model(&doc).Association("Tags").Clear(); err != nil {
		return err
	}

	if err := r.DB.Delete(&doc).Error; err != nil {
		return err
	}

	return nil
}
