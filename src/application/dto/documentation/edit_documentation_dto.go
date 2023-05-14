package documentationDTO

type EditDocumentationInputDTO struct {
	Title   string   `json:"title" validate:"required"`
	Content string   `json:"content" validate:"required"`
	Tags    []string `json:"tags"`
}

type EditDocumentationOutputDTO ListDocumentationsOutputDTO
