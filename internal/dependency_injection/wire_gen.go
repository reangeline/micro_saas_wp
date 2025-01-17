// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package dependency_injection

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/wire"
	"github.com/reangeline/micro_saas/internal/domain/contract/repository"
	usecase2 "github.com/reangeline/micro_saas/internal/domain/contract/usecasse"
	"github.com/reangeline/micro_saas/internal/domain/usecase"
	"github.com/reangeline/micro_saas/internal/infra/database"
	"github.com/reangeline/micro_saas/internal/presentation/controller"
)

// Injectors from wire.go:

func InitializeUser(vc *dynamodb.DynamoDB) (*controller.UserController, error) {
	userRepository := database.NewUserRepository(vc)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)
	return userController, nil
}

func InitializeCreateMessageWhatsApp() (*controller.WhatsAppController, error) {
	createMessageWhatsAppUseCase := usecase.NewCreateMessageWhatsAppUseCase()
	whatsAppController := controller.NewWhatsAppController(createMessageWhatsAppUseCase)
	return whatsAppController, nil
}

// wire.go:

var setUserUseCaseDependency = wire.NewSet(usecase.NewUserUseCase, wire.Bind(new(usecase2.UserUseCaseInterface), new(*usecase.UserUseCase)))

var setUserRepositoryDependency = wire.NewSet(database.NewUserRepository, wire.Bind(new(repository.UserRepositoryInterface), new(*database.UserRepository)))

var setCreateMessageWhatsAppUseCaseDependency = wire.NewSet(usecase.NewCreateMessageWhatsAppUseCase, wire.Bind(new(usecase2.CreateMessageWhatsAppUseCaseInterface), new(*usecase.CreateMessageWhatsAppUseCase)))
