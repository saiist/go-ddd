package users_test

import (
	app_service "go-ddd/src/application/users"
	domain_models "go-ddd/src/domain/models/users"
	"go-ddd/src/infrastructure/factory"
	"go-ddd/src/infrastructure/repositories/inmemory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRegisterService_Handle(t *testing.T) {
	inmemoryRepo := inmemory.NewUserRepository()
	service := app_service.NewUserRegisterService(
		factory.NewUserFactory(),
		inmemoryRepo,
		domain_models.NewUserService(inmemoryRepo),
	)

	// テストケース: ユーザー名が最小文字数の場合、成功するべき
	t.Run("success with min name", func(t *testing.T) {
		userName := "tes"
		err := service.Handle(userName)
		assert.NoError(t, err)

		users, err := inmemoryRepo.FindAll()
		assert.NoError(t, err)

		var result *domain_models.User
		for _, v := range users {
			if v.UserName.Value == userName {
				result = v
			}
		}

		assert.Equal(t, userName, result.UserName.Value)
	})

	// テストケース: ユーザー名が最大文字数の場合、成功するべき
	t.Run("success with long name", func(t *testing.T) {
		userName := "12345678901234567890"
		err := service.Handle(userName)
		assert.NoError(t, err)

		users, err := inmemoryRepo.FindAll()
		assert.NoError(t, err)

		var result *domain_models.User
		for _, v := range users {
			if v.UserName.Value == userName {
				result = v
			}
		}

		assert.Equal(t, userName, result.UserName.Value)
	})

	// テストケース: ユーザー名が空の場合、エラーが返されるべき
	t.Run("error empty name", func(t *testing.T) {
		err := service.Handle("")
		assert.Error(t, err)
	})

	// テストケース: ユーザーがすでに存在する場合、エラーが返されるべき
	t.Run("error exists", func(t *testing.T) {

		userName := "test"

		err := service.Handle(userName)
		assert.NoError(t, err)

		err = service.Handle(userName)
		assert.Error(t, err)
	})

	// テストケース: ユーザー名が短すぎる場合、エラーが返されるべき
	t.Run("error shot name", func(t *testing.T) {
		userName := "t"

		err := service.Handle(userName)
		assert.Error(t, err)
	})

	// テストケース: ユーザー名が長すぎる場合、エラーが返されるべき
	t.Run("error long name", func(t *testing.T) {
		userName := "123456789012345678901"

		err := service.Handle(userName)
		assert.Error(t, err)
	})

}
