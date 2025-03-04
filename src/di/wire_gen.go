// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"go-ddd/src/application/users"
	users2 "go-ddd/src/domain/models/users"
	"go-ddd/src/handler"
	"go-ddd/src/infrastructure/factory"
	"go-ddd/src/infrastructure/repositories"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeUserHandler(db *gorm.DB) *handler.UserHandler {
	iUserRepository := repositories.NewUserRepository(db)
	userGetService := users.NewUserGetService(iUserRepository)
	iUserFactory := factory.NewUserFactory()
	userService := users2.NewUserService(iUserRepository)
	userRegisterService := users.NewUserRegisterService(iUserFactory, iUserRepository, userService)
	userUpdateService := users.NewUserUpdateService(iUserRepository, userService)
	userDeleteService := users.NewUserDeleteService(iUserRepository)
	userHandler := handler.NewUserHandler(userGetService, userRegisterService, userUpdateService, userDeleteService)
	return userHandler
}

// wire.go:

var UserSet = wire.NewSet(repositories.NewUserRepository, factory.NewUserFactory, users2.NewUserService, users.NewUserGetService, users.NewUserRegisterService, users.NewUserUpdateService, users.NewUserDeleteService, handler.NewUserHandler)
