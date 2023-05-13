package documentationDTO

import (
	"time"

	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

type ListDocumentationsOutputDTO struct {
	Id        uint                           `json:"id"`
	Title     string                         `json:"title"`
	Content   string                         `json:"content"`
	CreatedAt time.Time                      `json:"created_at"`
	UpdatedAt time.Time                      `json:"updated_at"`
	Author    userDTO.UserWithoutPasswordDTO `json:"author"`
	Tags      []entity.Tag                   `json:"tags"`
}
