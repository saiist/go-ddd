package factory

import (
	domain_models "go-ddd/src/domain/models/users"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserFactory struct {
	db *gorm.DB
}

var _ domain_models.IUserFactory = &UserFactory{}

func NewUserFactory() domain_models.IUserFactory {
	return &UserFactory{}
}

func (f *UserFactory) Create(name string) (*domain_models.User, error) {
	id := uuid.New().String() // this is a dummy value

	user, err := domain_models.NewUser(domain_models.UserCreateConfig{Id: id, Name: name})
	if err != nil {
		return nil, err
	}

	return user, nil
}
