package tagRepository

import (
	tagDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/tag"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
	"gorm.io/gorm"
)

func (r *TagRepository) FindTagByName(tagName string) (*tagDTO.FindTagByNameOutput, error) {
	tag := entity.Tag{}

	if err := r.DB.First(&tag, "name = ?", tagName).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &tagDTO.FindTagByNameOutput{
		Id:   tag.Id,
		Name: tag.Name,
	}, nil
}
