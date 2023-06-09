package main

import (
	"auth-with-clean-architecture/dto"
	"auth-with-clean-architecture/internal/auth"
	"auth-with-clean-architecture/internal/customer"
	"auth-with-clean-architecture/internal/user"
	"auth-with-clean-architecture/pkg/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	godotenv.Load(".env")

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?parseTime=true"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal("initDB:", err)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, dto.Response{
			Meta: dto.MetaResponse{
				Success: true,
				Message: "Made with love by Alvriyanto Azis",
			},
			Data: nil,
		})
	})

	authHandler := auth.DefaultRequestHandler(db)
	userHandler := user.DefaultRequestHandler(db)

	r.POST("/auth/login", authHandler.Login)
	r.POST("/auth/register", userHandler.Create)

	authR := r.Group("/").Use(middleware.AuthMiddleware)
	authR.GET("/auth/profile", authHandler.ShowProfile)

	customerRepository := customer.NewRepository(db)
	customerUseCase := customer.NewUseCase(customerRepository)
	customerController := customer.NewController(customerUseCase)
	customerHandler := customer.NewRequestHandler(customerController)
	authR.GET("/customers", customerHandler.ShowAll)
	authR.POST("/customers", customerHandler.Create)
	authR.GET("/customers/:ID", customerHandler.Show)
	authR.PUT("/customers/:ID", customerHandler.Update)
	authR.DELETE("/customers/:ID", customerHandler.Destroy)

	authR.GET("/users", userHandler.ShowAll)
	authR.POST("/users", userHandler.Create)
	authR.GET("/users/:ID", userHandler.Show)
	authR.PUT("/users/:ID", userHandler.Update)
	authR.DELETE("/users/:ID", userHandler.Destroy)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
