package users

import (
	"go-ddd/src/domain/models/users"
)

type UserRegisterService struct {
	UserFactory    users.IUserFactory
	UserRepository users.IUserRepository
	UserService    *users.UserService
}

func NewUserRegisterService(
	userFactory users.IUserFactory,
	userRepository users.IUserRepository,
	userService *users.UserService,
) *UserRegisterService {
	return &UserRegisterService{
		UserFactory:    userFactory,
		UserRepository: userRepository,
		UserService:    userService,
	}
}

func (u *UserRegisterService) Handle(name string) error {

	user, err := u.UserFactory.Create(name)
	if err != nil {
		return err
	}

	if err := u.checkUserExists(user); err != nil {
		return err
	}

	return u.UserRepository.Save(user)
}

func (u *UserRegisterService) checkUserExists(user *users.User) error {
	exists, err := u.UserService.Exists(*user)
	if err != nil {
		return err
	}

	if exists {
		return &UserAlreadyExistsError{User: *user}
	}

	return nil
}
