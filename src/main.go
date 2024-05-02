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

func Main() {

	db := startDatabase()

	userHandler := di.InitializeUserHandler(db)

	r := gin.Default()
	users := r.Group("/users")
	{
		users.GET("/:id", userHandler.Get)
		users.POST("/", userHandler.Post)
		users.PUT("/:id", userHandler.Put)
		users.DELETE("/:id", userHandler.Delete)
	}

	startServer(r)
}

func startDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(getConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}
	return db
}

func startServer(r *gin.Engine) {
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
