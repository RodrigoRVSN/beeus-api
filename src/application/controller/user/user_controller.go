package userController

import (
	userUseCase "github.com/rodrigoRVSN/beeus-api/src/application/use_case/user"
)

type UserController struct {
	useCase userUseCase.UserUseCase
}

func NewUserController(useCase userUseCase.UserUseCase) *UserController {
	return &UserController{
		useCase: useCase,
	}
}
