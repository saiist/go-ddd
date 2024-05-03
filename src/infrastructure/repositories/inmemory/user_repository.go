package inmemory

import (
	domain_models "go-ddd/src/domain/models/users"
	"sync"
)

// this is a simple in-memory Store for users
var store = make(map[string]*domain_models.User)

type UserRepository struct {
	Store map[string]*domain_models.User
	mu    sync.RWMutex
}

var _ domain_models.IUserRepository = &UserRepository{}

func NewUserRepository() domain_models.IUserRepository {
	return &UserRepository{Store: store}
}

func (r *UserRepository) FindByName(name *domain_models.UserName) (*domain_models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, b := range r.Store {
		if b.UserName.Value == name.Value {
			// Create a deep copy of the user
			copyUser := *b
			return &copyUser, nil
		}
	}

	return nil, nil
}

func (r *UserRepository) FindById(id *domain_models.UserId) (*domain_models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result, ok := r.Store[id.Value]
	if !ok {
		return nil, nil
	}

	copyUser := *result
	return &copyUser, nil
}

func (r *UserRepository) FindAll() ([]*domain_models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var users []*domain_models.User
	for _, b := range r.Store {
		copyUser := *b
		users = append(users, &copyUser)
	}

	return users, nil
}

func (r *UserRepository) Save(user *domain_models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	copyUser := *user
	r.Store[copyUser.UserId.Value] = &copyUser
	return nil
}

func (r *UserRepository) Delete(user *domain_models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.Store, user.UserId.Value)
	return nil
}
