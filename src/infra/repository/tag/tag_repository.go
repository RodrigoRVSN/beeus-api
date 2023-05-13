package tagRepository

import (
	"github.com/rodrigoRVSN/beeus-api/src/domain/gateway"
	"gorm.io/gorm"
)

type TagRepository struct {
	DB *gorm.DB
}

func NewTagRepository(db *gorm.DB) gateway.TagGateway {
	return &TagRepository{DB: db}
}
