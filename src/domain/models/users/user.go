package domain_models

import "errors"

type User struct {
	UserId   UserId
	UserName UserName
}

type UserCreateConfig struct {
	Id   string
	Name string
}

type UserId string
type UserName string

func NewUser(conf UserCreateConfig) (*User, error) {
	uid, err := NewUserId(conf.Id)
	if err != nil {
		return nil, err
	}

	name, err := NewUserName(conf.Name)
	if err != nil {
		return nil, err
	}

	return &User{UserId: *uid, UserName: *name}, nil
}

type UserUpdateConfig struct {
	Name *string
}

func (u *User) Update(conf *UserUpdateConfig) error {
	if conf.Name != nil {
		name, err := NewUserName(*conf.Name)
		if err != nil {
			return err
		}

		u.UserName = *name
	}

	return nil
}

func (u *User) Equals(other *User) bool {
	return u.UserId == other.UserId
}

func NewUserId(value string) (*UserId, error) {
	if value == "" {
		return nil, errors.New("user id cannot be empty")
	}

	userId := UserId(value)
	return &userId, nil

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

	userName := UserName(value)
	return &userName, nil
}

type IUserRepository interface {
	FindAll() ([]*User, error)
	FindByName(name *UserName) (*User, error)
	FindById(id *UserId) (*User, error)
	Save(user *User) error
	Delete(user *User) error
}

type IUserFactory interface {
	Create(name string) (*User, error)
}
