package main

import (
	"go-crud-api/database"
	_ "go-crud-api/docs" // This is required for swagger docs
	"go-crud-api/handlers"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Example API
// @version 1.0
// @description This is a simple CRUD API with Go and Gin.
// @termsOfService https://example.com/terms/
// @contact.name API Support
// @contact.url https://example.com/support
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/users", handlers.GetUsers(db))
	r.GET("/users/:id", handlers.GetUser(db))
	r.POST("/users", handlers.CreateUser(db))
	r.PUT("/users/:id", handlers.UpdateUser(db))
	r.DELETE("/users/:id", handlers.DeleteUser(db))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Server Error: ", err)
	}

}
