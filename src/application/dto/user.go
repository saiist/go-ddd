package dto

import "go-ddd/src/domain/entity"

type UserData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewUserData(source entity.User) *UserData {
	return &UserData{
		Id:   source.UserId.Value,
		Name: source.UserName.Value,
	}
}
