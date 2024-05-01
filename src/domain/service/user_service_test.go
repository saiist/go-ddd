package domain_service

import (
	"go-ddd/src/domain/entity"
	"go-ddd/src/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_Exists(t *testing.T) {
	mockRepo := mocks.NewIUserRepository(t)

	user, err := entity.NewUser("John Doe")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Set up the expectation
	mockRepo.On("Find", &user.UserName).Return(nil, nil).Once()

	// Create the service with the mock
	service := NewUserService(mockRepo)

	// Call the method and check the result
	exists, err := service.Exists(*user)
	assert.False(t, exists)
	assert.NoError(t, err)

}
