package main

import (
	"go-crud-api/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r := gin.Default()

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Server Error: ", err)
	}

}
