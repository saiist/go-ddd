package users

import (
	"errors"
)

type UserId struct {
	Value string
}

func NewUserId(value string) (*UserId, error) {
	if value == "" {
		return nil, errors.New("user id cannot be empty")
	}

	return &UserId{Value: value}, nil
}
