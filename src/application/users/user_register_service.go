package users

import (
	"go-ddd/src/domain/models/users"
)

type UserRegisterService struct {
	UserRepository users.IUserRepository
	UserService    *users.UserService
}

func NewUserRegisterService(
	userRepository users.IUserRepository,
	userService *users.UserService,
) *UserRegisterService {
	return &UserRegisterService{
		UserRepository: userRepository,
		UserService:    userService,
	}
}

func (u *UserRegisterService) Handle(name string) error {
	user, err := users.NewUser(users.UserCreateConfig{Name: name})
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
