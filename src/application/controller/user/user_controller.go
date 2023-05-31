package userController

import (
	userUseCase "github.com/rodrigoRVSN/beeus-api/src/application/use_case/user"
	"github.com/rodrigoRVSN/beeus-api/src/infra/context"
)

type UserControllerInterface interface {
	CreateUser(ctx context.Context) error
	SignInUser(ctx context.Context) error
	Me(ctx context.Context) error
	GetRankedUsers(ctx context.Context) error
}

type UserController struct {
	useCase userUseCase.UserControllerInterface
}

func NewUserController(useCase userUseCase.UserControllerInterface) UserControllerInterface {
	return &UserController{
		useCase: useCase,
	}
}
