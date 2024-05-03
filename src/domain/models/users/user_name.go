package domain_models

import (
	"errors"
)

type UserName struct {
	Value string
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
