package services

import (
	"go-ddd/src/domain/models/users"
)

type UserService struct {
	UserRepository users.IUserRepository
}

func NewUserService(userRepository users.IUserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (u *UserService) Exists(user users.User) (bool, error) {
	found, err := u.UserRepository.FindByName(&user.UserName)
	if err != nil {
		return false, err
	}

	return found != nil, nil
}
