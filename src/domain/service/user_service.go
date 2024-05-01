package domain_service

import (
	"go-ddd/src/domain/entity"
	repo "go-ddd/src/repository"
)

type UserService struct {
	UserRepository repo.IUserRepository
}

func NewUserService(userRepository repo.IUserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (u *UserService) Exists(user entity.User) (bool, error) {
	found, err := u.UserRepository.FindByName(&user.UserName)
	if err != nil {
		return false, err
	}

	return found != nil, nil
}
