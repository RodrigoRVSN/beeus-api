package userDtos

type SignInInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInOutput struct {
	Token  string `json:"access_token"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}