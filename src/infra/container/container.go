package container

import (
	documentationController "github.com/rodrigoRVSN/beeus-api/src/application/controller/documentation"
	userController "github.com/rodrigoRVSN/beeus-api/src/application/controller/user"
	documentationUseCase "github.com/rodrigoRVSN/beeus-api/src/application/use_case/documentation"
	userUseCase "github.com/rodrigoRVSN/beeus-api/src/application/use_case/user"
	"github.com/rodrigoRVSN/beeus-api/src/config"
	documentationRepository "github.com/rodrigoRVSN/beeus-api/src/infra/repository/documentation"
	tagRepository "github.com/rodrigoRVSN/beeus-api/src/infra/repository/tag"
	userRepository "github.com/rodrigoRVSN/beeus-api/src/infra/repository/user"
)

type Container struct {
	DocumentationController documentationController.DocumentationControllerInterface
	UserController          userController.UserControllerInterface
}

func InitializeContainer() *Container {
	db := config.ConnectDb()

	tagRepository := tagRepository.NewTagRepository(db)

	userRepository := userRepository.NewUserRepository(db)
	userUseCase := userUseCase.NewUserUseCase(userRepository)
	userController := userController.NewUserController(userUseCase)

	documentationRepository := documentationRepository.NewDocumentationRepository(db)
	documentationUseCase := documentationUseCase.NewDocumentationUseCase(documentationRepository, userRepository, tagRepository)
	documentationController := documentationController.NewDocumentationController(documentationUseCase)

	return &Container{
		DocumentationController: documentationController,
		UserController:          userController,
	}
}
