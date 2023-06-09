package documentationDTO

type CreateDocumentationInputDTO struct {
	Title   string   `json:"title" validate:"required"`
	Content string   `json:"content" validate:"required"`
	Tags    []string `json:"tags"`
}

type CreateDocumentationOutputDTO ListDocumentationsOutputDTO
