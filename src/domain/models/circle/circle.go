package domain_models

import (
	domain_models_users "go-ddd/src/domain/models/users"
)

type Circle struct {
	Id   domain_models_users.UserId
	Name CircleName
}

type ICircleFactory interface {
	Create(id domain_models_users.UserId, name string) (*Circle, error)
}
