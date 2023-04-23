package userController

import (
	userUseCase "github.com/rodrigoRVSN/beeus-api/src/application/use_case/user"
)

type UserController struct {
	useCase userUseCase.CreateUserUseCase
}

func NewUserController(useCase userUseCase.CreateUserUseCase) *UserController {
	return &UserController{
		useCase: useCase,
	}
}
