package users

import (
	"errors"
	"go-ddd/src/application/dtos"
	"go-ddd/src/domain/models/users"
)

type UserGetService struct {
	UserRepository users.IUserRepository
}

func NewUserGetService(
	userRepository users.IUserRepository,
) *UserGetService {
	return &UserGetService{
		UserRepository: userRepository,
	}
}

func (u *UserGetService) Get(id string) (*dtos.UserData, error) {
	user, err := u.findUserById(id)
	if err != nil {
		if errors.Is(err, &UserNotFoundError{}) {
			return nil, nil
		}

		return nil, err
	}

	return dtos.NewUserData(*user), nil
}

func (u *UserGetService) GetAll() ([]*dtos.UserData, error) {
	users, err := u.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var result []*dtos.UserData
	for _, user := range users {
		result = append(result, dtos.NewUserData(*user))
	}

	return result, nil
}

func (u *UserGetService) findUserById(id string) (*users.User, error) {
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
