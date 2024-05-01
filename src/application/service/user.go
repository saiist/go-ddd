package app_service

import (
	"fmt"
	"go-ddd/src/application/dto"
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

func (u *UserAppService) Get(id string) (*dto.UserData, error) {
	targetId, err := entity.NewUserId(id)
	if err != nil {
		return nil, err
	}

	user, err := u.UserRepository.FindById(targetId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return dto.NewUserData(*user), nil
}
