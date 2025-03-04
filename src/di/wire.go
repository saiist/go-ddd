//go:build wireinject
// +build wireinject

package di

import (
	app_service_users "go-ddd/src/application/users"
	domain_models "go-ddd/src/domain/models/users"
	"go-ddd/src/handler"
	"go-ddd/src/infrastructure/factory"
	"go-ddd/src/infrastructure/repositories"

	// repo "go-ddd/src/infrastructure/repositories/inmemory"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var UserSet = wire.NewSet(
	repositories.NewUserRepository,
	factory.NewUserFactory,
	domain_models.NewUserService,
	app_service_users.NewUserGetService,
	app_service_users.NewUserRegisterService,
	app_service_users.NewUserUpdateService,
	app_service_users.NewUserDeleteService,
	handler.NewUserHandler,
)

func InitializeUserHandler(db *gorm.DB) *handler.UserHandler {
	wire.Build(UserSet)
	return nil
}
