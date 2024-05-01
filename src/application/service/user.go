package app_service

import (
	"fmt"
	"go-ddd/src/domain/entity"
	domain_service "go-ddd/src/domain/service"
	"go-ddd/src/repository"
)

type UserAppService struct {
	UserRepository repository.IUserRepository
	UserService    *domain_service.UserService
}

func NewUserAppService(
	userRepository repository.IUserRepository,
	userService *domain_service.UserService,
) *UserAppService {
	return &UserAppService{
		UserRepository: userRepository,
		UserService:    userService,
	}
}

func (u *UserAppService) Register(name string) error {
	user, err := entity.NewUser(name)
	if err != nil {
		return err
	}

	exists, err := u.UserService.Exists(*user)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("error checking if user exists: %s", name)
	}

	err = u.UserRepository.Save(user)
	if err != nil {
		return err
	}

	return nil
}
