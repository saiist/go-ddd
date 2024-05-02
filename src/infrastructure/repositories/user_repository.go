package repositories

import (
	"errors"
	"go-ddd/src/domain/models/users"
	"go-ddd/src/infrastructure/data_models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

var _ users.IUserRepository = &UserRepository{}

func NewUserRepository(db *gorm.DB) users.IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByName(name *users.UserName) (*users.User, error) {
	var model data_models.UserDataModel
	result := r.db.Where("name = ?", name.Value).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return model.ToEntity()
}

func (r *UserRepository) FindById(id *users.UserId) (*users.User, error) {
	var model data_models.UserDataModel
	result := r.db.Where("id = ?", id.Value).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return model.ToEntity()
}

func (r *UserRepository) FindAll() ([]*users.User, error) {
	var models []data_models.UserDataModel
	result := r.db.Find(&models)
	if result.Error != nil {
		return nil, result.Error
	}

	var users []*users.User
	for _, model := range models {
		user, err := model.ToEntity()
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) Save(user *users.User) error {
	model := data_models.UserDataModel{}.ToDataModel(user)
	result := r.db.Save(&model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) Delete(user *users.User) error {
	model := data_models.UserDataModel{}.ToDataModel(user)
	result := r.db.Delete(&model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
