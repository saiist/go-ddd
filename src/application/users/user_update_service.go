package users

import (
	domain_models "go-ddd/src/domain/models/users"
)

type UserUpdateService struct {
	UserRepository domain_models.IUserRepository
	UserService    *domain_models.UserService
}

func NewUserUpdateService(
	userRepository domain_models.IUserRepository,
	userService *domain_models.UserService,
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

	conf := domain_models.UserUpdateConfig{Name: &name}
	if err := user.Update(&conf); err != nil {
		return err
	}

	if err := u.checkUserExists(user); err != nil {
		return err
	}

	return u.UserRepository.Save(user)
}

func (u *UserUpdateService) findUserById(id string) (*domain_models.User, error) {
	targetId, err := domain_models.NewUserId(id)
	if err != nil {
		return nil, err
	}

	user, err := u.UserRepository.FindById(targetId)
	if err != nil {
		return nil, &UserNotFoundError{Id: *targetId}
	}

	return user, nil
}

func (u *UserUpdateService) checkUserExists(user *domain_models.User) error {
	exists, err := u.UserService.Exists(*user)
	if err != nil {
		return err
	}

	if exists {
		return &UserAlreadyExistsError{User: *user}
	}

	return nil
}
