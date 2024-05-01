package data_models

import (
	"fmt"
	"go-ddd/src/domain/models/users"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserDataModel struct {
	Id   string `gorm:"primaryKey"`
	Name string `validate:"required,min=3"`
}

func (d *UserDataModel) ToEntity() (*users.User, error) {
	uid, err := users.NewUserId(d.Id)
	if err != nil {
		return nil, err
	}

	un, err := users.NewUserName(d.Name)
	if err != nil {
		return nil, err
	}

	return &users.User{UserId: *uid, UserName: *un}, nil
}

func (UserDataModel) ToDataModel(from *users.User) *UserDataModel {
	return &UserDataModel{
		Id:   from.UserId.Value,
		Name: from.UserName.Value,
	}
}

func (UserDataModel) TableName() string {
	return "users"
}

func (u *UserDataModel) BeforeCreate(tx *gorm.DB) (err error) {
	if err := u.validateDataModel(); err != nil {
		return err
	}
	return
}

func (u *UserDataModel) BeforeUpdate(tx *gorm.DB) (err error) {
	if err := u.validateDataModel(); err != nil {
		return err
	}
	return
}

func (u *UserDataModel) validateDataModel() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			return fmt.Errorf("validation failed on '%s': condition '%s' not met", field, tag)
		}
	}
	return nil
}
