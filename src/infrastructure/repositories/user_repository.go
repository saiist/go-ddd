package repositories

import (
	"errors"
	domain_models "go-ddd/src/domain/models/users"
	"go-ddd/src/infrastructure/data_models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

var _ domain_models.IUserRepository = &UserRepository{}

func NewUserRepository(db *gorm.DB) domain_models.IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByName(name *domain_models.UserName) (*domain_models.User, error) {
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

func (r *UserRepository) FindById(id *domain_models.UserId) (*domain_models.User, error) {
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

func (r *UserRepository) FindAll() ([]*domain_models.User, error) {
	var models []data_models.UserDataModel
	result := r.db.Find(&models)
	if result.Error != nil {
		return nil, result.Error
	}

	var users []*domain_models.User
	for _, model := range models {
		user, err := model.ToEntity()
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) Save(user *domain_models.User) error {
	model := data_models.UserDataModel{}.ToDataModel(user)
	result := r.db.Save(&model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) Delete(user *domain_models.User) error {
	model := data_models.UserDataModel{}.ToDataModel(user)
	result := r.db.Delete(&model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
