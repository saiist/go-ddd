package data_models

import (
	"fmt"
	domain_models "go-ddd/src/domain/models/users"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserDataModel struct {
	Id   string `gorm:"primaryKey"`
	Name string `validate:"required,min=3"`
}

func (d *UserDataModel) ToEntity() (*domain_models.User, error) {
	uid, err := domain_models.NewUserId(d.Id)
	if err != nil {
		return nil, err
	}

	un, err := domain_models.NewUserName(d.Name)
	if err != nil {
		return nil, err
	}

	return &domain_models.User{UserId: *uid, UserName: *un}, nil
}

func (UserDataModel) ToDataModel(from *domain_models.User) *UserDataModel {
	return &UserDataModel{
		Id:   string(from.UserId),
		Name: string(from.UserName),
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
