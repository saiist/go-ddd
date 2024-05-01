package users_test

import (
	"go-ddd/src/domain/models/users"
	"go-ddd/src/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_Exists(t *testing.T) {
	mockRepo := mocks.NewIUserRepository(t)

	user, err := users.NewUser(users.UserCreateConfig{Name: "Test"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Set up the expectation
	mockRepo.On("FindByName", &user.UserName).Return(nil, nil).Once()

	// Create the service with the mock
	service := users.NewUserService(mockRepo)

	// Call the method and check the result
	exists, err := service.Exists(*user)
	assert.False(t, exists)
	assert.NoError(t, err)

}
