package main

import (
	"go-crud-api/database"
	"go-crud-api/handlers"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Server Error: ", err)
	}

}
