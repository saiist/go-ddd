package src

import (
	"fmt"
	"log"
	"os"

	"go-ddd/src/di"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
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

	userHandler := di.InitializeUserHandler(db)

	r := gin.Default()
	r.GET("/users/:id", userHandler.Get)
	r.POST("/users", userHandler.Post)
	r.PUT("/users/:id", userHandler.Put)
	r.DELETE("/users/:id", userHandler.Delete)

	err := r.Run(":3110")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func getConnectionString() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_HOSTNAME"),
	)
}
