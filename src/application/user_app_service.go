package application

import (
	"errors"
	"go-ddd/src/application/dtos"
	"go-ddd/src/domain/models/users"
)

type UserAppService struct {
	UserRepository users.IUserRepository
	UserService    *users.UserService
}

func NewUserAppService(
	userRepository users.IUserRepository,
	userService *users.UserService,
) *UserAppService {
	return &UserAppService{
		UserRepository: userRepository,
		UserService:    userService,
	}
}

func (u *UserAppService) Register(name string) error {
	user, err := users.NewUser(name)
	if err != nil {
		return err
	}

	if err := u.checkUserExists(user); err != nil {
		return err
	}

	return u.UserRepository.Save(user)
}

func (u *UserAppService) Get(id string) (*dtos.UserData, error) {
	user, err := u.findUserById(id)
	if err != nil {
		if errors.Is(err, &UserNotFoundError{}) {
			return nil, nil
		}

		return nil, err
	}

	return dtos.NewUserData(*user), nil
}

func (u *UserAppService) Update(id string, name string) error {
	user, err := u.findUserById(id)
	if err != nil {
		return err
	}

	newName, err := users.NewUserName(name)
	if err != nil {
		return err
	}

	err = user.ChangeName(*newName)
	if err != nil {
		return err
	}

	if err := u.checkUserExists(user); err != nil {
		return err
	}

	return u.UserRepository.Save(user)
}

func (u *UserAppService) findUserById(id string) (*users.User, error) {
	targetId, err := users.NewUserId(id)
	if err != nil {
		return nil, err
	}

	user, err := u.UserRepository.FindById(targetId)
	if err != nil {
		return nil, &UserNotFoundError{Id: *targetId}
	}

	return user, nil
}

func (u *UserAppService) checkUserExists(user *users.User) error {
	exists, err := u.UserService.Exists(*user)
	if err != nil {
		return err
	}

	if exists {
		return &UserAlreadyExistsError{User: *user}
	}

	return nil
}
