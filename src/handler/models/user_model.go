package models

import "go-ddd/src/application/dtos"

type UserResponseModel struct {
	Id   string
	Name string
}

func NewUserResponseModel(source *dtos.UserData) *UserResponseModel {
	return &UserResponseModel{
		Id:   source.Id,
		Name: source.Name,
	}
}

type UserGetResponseModel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewUserGetResponseModel(user *UserResponseModel) *UserGetResponseModel {
	return &UserGetResponseModel{
		Id:   user.Id,
		Name: user.Name,
	}
}

type UserIndexResponseModel struct {
	Users []*UserResponseModel
}

func NewUserIndexResponseModel(users []*dtos.UserData) *UserIndexResponseModel {
	var result []*UserResponseModel
	for _, user := range users {
		result = append(result, NewUserResponseModel(user))
	}

	return &UserIndexResponseModel{Users: result}
}

type UserPostRequestModel struct {
	Name string `json:"name"`
}

type UserPutRequestModel struct {
	Name string `json:"name"`
}
