package dtos

import "go-ddd/src/domain/models/users"

type UserData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewUserData(source users.User) *UserData {
	return &UserData{
		Id:   source.UserId.Value,
		Name: source.UserName.Value,
	}
}
