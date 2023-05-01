package userDTO

type SignInInputDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignInOutputDTO struct {
	Token string `json:"access_token"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
