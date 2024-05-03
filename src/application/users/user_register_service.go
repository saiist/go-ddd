package users

import (
	domain_models "go-ddd/src/domain/models/users"
)

type UserRegisterService struct {
	UserFactory    domain_models.IUserFactory
	UserRepository domain_models.IUserRepository
	UserService    *domain_models.UserService
}

func NewUserRegisterService(
	userFactory domain_models.IUserFactory,
	userRepository domain_models.IUserRepository,
	userService *domain_models.UserService,
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

func (u *UserRegisterService) checkUserExists(user *domain_models.User) error {
	exists, err := u.UserService.Exists(*user)
	if err != nil {
		return err
	}

	if exists {
		return &UserAlreadyExistsError{User: *user}
	}

	return nil
}
