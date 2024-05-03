package domain_models

type User struct {
	UserId   UserId
	UserName UserName
}

type UserCreateConfig struct {
	Id   string
	Name string
}

func NewUser(conf UserCreateConfig) (*User, error) {
	uid, err := NewUserId(conf.Id)
	if err != nil {
		return nil, err
	}

	un, err := NewUserName(conf.Name)
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
	FindAll() ([]*User, error)
	FindByName(name *UserName) (*User, error)
	FindById(id *UserId) (*User, error)
	Save(user *User) error
	Delete(user *User) error
}

type IUserFactory interface {
	Create(name string) (*User, error)
}
