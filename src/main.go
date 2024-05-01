package src

import (
	"fmt"
	application_service "go-ddd/src/application/service"
	domain_service "go-ddd/src/domain/services"
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
	userService := domain_service.NewUserService(userRepository)
	userAppService := application_service.NewUserAppService(userRepository, userService)

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
