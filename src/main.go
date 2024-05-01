package src

import (
	"fmt"
	"go-ddd/src/application"
	"go-ddd/src/domain/models/users"
	repo "go-ddd/src/infrastructure/repositories"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Main() {
	err := CreateUser("John smith")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("User created")
}

func CreateUser(name string) error {
	db, err := gorm.Open(postgres.Open(getConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}

	userRepository := repo.NewUserRepository(db)
	userService := users.NewUserService(userRepository)
	userAppService := application.NewUserAppService(userRepository, userService)

	err = userAppService.Register(name)

	return err
}

func getConnectionString() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_HOSTNAME"),
	)
}
