package users

import (
	"go-ddd/src/domain/models/users"
)

type UserUpdateService struct {
	UserRepository users.IUserRepository
	UserService    *users.UserService
}

func NewUserUpdateService(
	userRepository users.IUserRepository,
	userService *users.UserService,
) *UserUpdateService {
	return &UserUpdateService{
		UserRepository: userRepository,
		UserService:    userService,
	}
}

func (u *UserUpdateService) Update(id string, name string) error {
	user, err := u.findUserById(id)
	if err != nil {
		return err
	}

	conf := users.UserUpdateConfig{Name: &name}
	if err := user.Update(&conf); err != nil {
		return err
	}

	if err := u.checkUserExists(user); err != nil {
		return err
	}

	return u.UserRepository.Save(user)
}

func (u *UserUpdateService) findUserById(id string) (*users.User, error) {
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

func (u *UserUpdateService) checkUserExists(user *users.User) error {
	exists, err := u.UserService.Exists(*user)
	if err != nil {
		return err
	}

	if exists {
		return &UserAlreadyExistsError{User: *user}
	}

	return nil
}
