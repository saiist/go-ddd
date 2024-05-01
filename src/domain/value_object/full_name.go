package valueobject

import (
	"errors"
	"unicode"
)

type FullName struct {
	FirstName string
	LastName  string
}

func NewFullName(firstName string, lastName string) (*FullName, error) {
	if firstName == "" {
		return nil, errors.New("first name cannot be empty")
	}
	if lastName == "" {
		return nil, errors.New("last name cannot be empty")
	}

	if !validateName(firstName) {
		return nil, errors.New("first name should contain only alphabets")
	}

	if !validateName(lastName) {
		return nil, errors.New("last name should contain only alphabets")
	}

	return &FullName{FirstName: firstName, LastName: lastName}, nil
}

func (f *FullName) GetFullName() string {
	return f.FirstName + " " + f.LastName
}

func (f *FullName) Equals(other *FullName) bool {
	return f.FirstName == other.FirstName && f.LastName == other.LastName
}

func validateName(value string) bool {
	// Restrict to alphabets only
	for _, char := range value {
		if !unicode.IsLetter(char) {
			return false
		}
	}

	return true
}
