package repository

import (
	"errors"
	"go-ddd/src/domain/entity"
	"go-ddd/src/infrastructure/data_model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindByName(name *entity.UserName) (*entity.User, error)
	FindById(id *entity.UserId) (*entity.User, error)
	Save(user *entity.User) error
	Delete(user *entity.User) error
}

type UserRepository struct {
	db *gorm.DB
}

var _ IUserRepository = &UserRepository{}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByName(name *entity.UserName) (*entity.User, error) {
	var model data_model.UserDataModel
	result := r.db.Where("name = ?", name.Value).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return model.ToEntity()
}

func (r *UserRepository) FindById(id *entity.UserId) (*entity.User, error) {
	var model data_model.UserDataModel
	result := r.db.Where("id = ?", id.Value).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return model.ToEntity()
}

func (r *UserRepository) Save(user *entity.User) error {
	model := data_model.UserDataModel{}.ToDataModel(user)
	result := r.db.Save(&model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) Delete(user *entity.User) error {
	model := data_model.UserDataModel{}.ToDataModel(user)
	result := r.db.Delete(&model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
