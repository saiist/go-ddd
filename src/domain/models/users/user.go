package users

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	UserId   UserId
	UserName UserName
}

type UserId struct {
	Value string
}

type UserName struct {
	Value string
}

func NewUser(name string) (*User, error) {
	un, err := NewUserName(name)
	if err != nil {
		return nil, err
	}

	uid, err := NewUserId(uuid.New().String())
	if err != nil {
		return nil, err
	}

	return &User{UserId: *uid, UserName: *un}, nil
}

func (u *User) Equals(other *User) bool {
	return u.UserId.Value == other.UserId.Value
}

func (u *User) ChangeName(name UserName) error {
	u.UserName = name
	return nil
}

func NewUserId(value string) (*UserId, error) {
	if value == "" {
		return nil, errors.New("user id cannot be empty")
	}

	return &UserId{Value: value}, nil
}

func NewUserName(value string) (*UserName, error) {
	if value == "" {
		return nil, errors.New("user name cannot be empty")
	}

	if len(value) < 3 {
		return nil, errors.New("user name should be at least 3 characters long")
	}

	if len(value) > 20 {
		return nil, errors.New("user name should not be more than 20 characters long")
	}

	return &UserName{Value: value}, nil
}

type IUserRepository interface {
	FindByName(name *UserName) (*User, error)
	FindById(id *UserId) (*User, error)
	Save(user *User) error
	Delete(user *User) error
}
