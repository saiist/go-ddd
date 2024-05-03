package domain_models

import "errors"

type CircleName struct {
	Value string
}

func NewCircleName(value string) (*CircleName, error) {
	if value == "" {
		return nil, errors.New("circle name cannot be empty")
	}

	return &CircleName{Value: value}, nil
}
