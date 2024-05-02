package src

import (
	"fmt"
	"log"
	"os"

	"go-ddd/src/di"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(postgres.Open(getConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}
}

func Main() {
	err := CreateUser("John smith")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("User created")
}

func CreateUser(name string) error {

	userRegisterService := di.InitializeUserRegisterService(db)
	err := userRegisterService.Handle(name)

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
