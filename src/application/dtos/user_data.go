package dtos

import domain_models "go-ddd/src/domain/models/users"

type UserData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewUserData(source domain_models.User) *UserData {
	return &UserData{
		Id:   string(source.UserId),
		Name: string(source.UserName),
	}
}
