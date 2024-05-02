package users

import (
	"errors"
	"go-ddd/src/domain/models/users"
)

type UserDeleteService struct {
	UserRepository users.IUserRepository
}

func NewUserDeleteService(
	userRepository users.IUserRepository,
) *UserDeleteService {
	return &UserDeleteService{
		UserRepository: userRepository,
	}
}

func (u *UserDeleteService) Handle(id string) error {
	user, err := u.findUserById(id)
	if err != nil {
		if errors.Is(err, &UserNotFoundError{}) {
			// If the user is not found, we consider the operation successful
			return nil
		}

		return err
	}

	return u.UserRepository.Delete(user)
}

func (u *UserDeleteService) findUserById(id string) (*users.User, error) {
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
