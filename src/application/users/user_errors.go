package users

import (
	"fmt"
	"go-ddd/src/domain/models/users"
)

type UserAlreadyExistsError struct {
	User users.User
}

func (e *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("user already exists Id: %s, Name: %s",
		e.User.UserId.Value, e.User.UserName.Value)
}

type UserNotFoundError struct {
	Id users.UserId
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user not found: %s", e.Id.Value)
}
