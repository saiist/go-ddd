package users

import (
	"github.com/google/uuid"
)

type User struct {
	UserId   UserId
	UserName UserName
}

type UserCreateConfig struct {
	Name string
}

func NewUser(conf UserCreateConfig) (*User, error) {
	un, err := NewUserName(conf.Name)
	if err != nil {
		return nil, err
	}

	uid, err := NewUserId(uuid.New().String())
	if err != nil {
		return nil, err
	}

	return &User{UserId: *uid, UserName: *un}, nil
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
	return u.UserId.Value == other.UserId.Value
}

type IUserRepository interface {
	FindByName(name *UserName) (*User, error)
	FindById(id *UserId) (*User, error)
	Save(user *User) error
	Delete(user *User) error
}
