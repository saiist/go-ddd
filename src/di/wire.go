//go:build wireinject
// +build wireinject

package di

import (
	app_service_users "go-ddd/src/application/users"
	models_users "go-ddd/src/domain/models/users"
	repo "go-ddd/src/infrastructure/repositories"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeUserRegisterService(db *gorm.DB) *app_service_users.UserRegisterService {
	wire.Build(repo.NewUserRepository, models_users.NewUserService, app_service_users.NewUserRegisterService)
	return &app_service_users.UserRegisterService{}
}
