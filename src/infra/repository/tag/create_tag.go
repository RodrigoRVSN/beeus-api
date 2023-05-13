package tagRepository

import (
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (r *TagRepository) CreateTag(tagName string) (*entity.Tag, error) {
	tag := entity.Tag{
		Name: tagName,
	}

	if err := r.DB.Create(&tag).Error; err != nil {
		return nil, err
	}

	return &tag, nil
}
