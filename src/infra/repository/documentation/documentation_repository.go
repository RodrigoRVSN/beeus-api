package documentationRepository

import (
	"github.com/rodrigoRVSN/beeus-api/src/domain/gateway"
	"gorm.io/gorm"
)

type DocumentationRepository struct {
	DB *gorm.DB
}

func NewDocumentationRepository(db *gorm.DB) gateway.DocumentationGateway {
	return &DocumentationRepository{DB: db}
}
