package factory

import (
	domain_models "go-ddd/src/domain/models/circle"
	domain_models_users "go-ddd/src/domain/models/users"
)

type CircleFactory struct {
}

var _ domain_models.ICircleFactory = &CircleFactory{}

func NewCircleFactory() domain_models.ICircleFactory {
	return &CircleFactory{}
}

func (f *CircleFactory) Create(id domain_models_users.UserId, name string) (*domain_models.Circle, error) {
	cname, err := domain_models.NewCircleName(name)
	if err != nil {
		return nil, err
	}

	return &domain_models.Circle{Id: id, Name: *cname}, nil
}
