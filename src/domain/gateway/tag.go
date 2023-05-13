package gateway

import (
	tagDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/tag"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

type TagGateway interface {
	CreateTag(tagName string) (*entity.Tag, error)
	FindTagByName(tagName string) (*tagDTO.FindTagByNameOutput, error)
}
