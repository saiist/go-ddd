package users_test

import (
	app_service "go-ddd/src/application/users"
	domain_models "go-ddd/src/domain/models/users"
	"go-ddd/src/infrastructure/repositories/inmemory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRegisterService_Handle(t *testing.T) {
	inmemoryRepo := inmemory.NewUserRepository()
	service := app_service.NewUserRegisterService(
		inmemoryRepo,
		domain_models.NewUserService(inmemoryRepo),
	)

	// テストケース: ユーザー名が最小文字数の場合、成功するべき
	t.Run("success with min name", func(t *testing.T) {
		user, err := domain_models.NewUser(domain_models.UserCreateConfig{Name: "tes"})
		assert.NoError(t, err)

		err = service.Handle(user.UserName.Value)
		assert.NoError(t, err)

		result, _ := service.UserRepository.FindByName(&user.UserName)
		assert.Equal(t, user.UserName.Value, result.UserName.Value)
	})

	// テストケース: ユーザー名が最大文字数の場合、成功するべき
	t.Run("success with long name", func(t *testing.T) {
		user, err := domain_models.NewUser(domain_models.UserCreateConfig{Name: "12345678901234567890"})
		assert.NoError(t, err)

		err = service.Handle(user.UserName.Value)
		assert.NoError(t, err)

		result, _ := service.UserRepository.FindByName(&user.UserName)
		assert.Equal(t, user.UserName.Value, result.UserName.Value)
	})

	// テストケース: ユーザー名が空の場合、エラーが返されるべき
	t.Run("error empty name", func(t *testing.T) {
		err := service.Handle("")
		assert.Error(t, err)
	})

	// テストケース: ユーザーがすでに存在する場合、エラーが返されるべき
	t.Run("error exists", func(t *testing.T) {
		user, _ := domain_models.NewUser(domain_models.UserCreateConfig{Name: "test"})

		err := service.Handle(user.UserName.Value)
		assert.NoError(t, err)

		err = service.Handle(user.UserName.Value)
		assert.Error(t, err)
	})

	// テストケース: ユーザー名が短すぎる場合、エラーが返されるべき
	t.Run("error shot name", func(t *testing.T) {
		_, err := domain_models.NewUser(domain_models.UserCreateConfig{Name: "t"})
		assert.Error(t, err)
	})

	// テストケース: ユーザー名が長すぎる場合、エラーが返されるべき
	t.Run("error long name", func(t *testing.T) {
		_, err := domain_models.NewUser(domain_models.UserCreateConfig{Name: "123456789012345678901"})
		assert.Error(t, err)
	})

}
