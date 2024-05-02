package inmemory

import (
	"go-ddd/src/domain/models/users"
	"sync"

	"gorm.io/gorm"
)

// this is a simple in-memory store for users
var store = make(map[string]*users.User)

type UserRepository struct {
	mu sync.RWMutex
}

var _ users.IUserRepository = &UserRepository{}

func NewUserRepository(*gorm.DB) users.IUserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByName(name *users.UserName) (*users.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, b := range store {
		if b.UserName.Value == name.Value {
			// Create a deep copy of the user
			copyUser := *b
			return &copyUser, nil
		}
	}

	return nil, nil
}

func (r *UserRepository) FindById(id *users.UserId) (*users.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result, ok := store[id.Value]
	if !ok {
		return nil, nil
	}

	copyUser := *result
	return &copyUser, nil
}

func (r *UserRepository) Save(user *users.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	copyUser := *user
	store[copyUser.UserId.Value] = &copyUser
	return nil
}

func (r *UserRepository) Delete(user *users.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(store, user.UserId.Value)
	return nil
}
