package users

import (
	"fmt"
	domain_models "go-ddd/src/domain/models/users"
)

type UserAlreadyExistsError struct {
	User domain_models.User
}

func (e *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("user already exists Id: %s, Name: %s",
		e.User.UserId, e.User.UserName)
}

type UserNotFoundError struct {
	Id domain_models.UserId
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user not found: %s", e.Id)
}
