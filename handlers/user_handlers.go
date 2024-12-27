package handlers

import (
	"database/sql"
	"go-crud-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, name, email, created_at FROM users")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		defer rows.Close()
		var users []models.User

		for rows.Next() {
			var user models.User
			if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			users = append(users, user)
		}

		c.JSON(http.StatusOK, users)
	}
}

func GetUser(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.User
		err := db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)

	}
}

func CreateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User
		// Creates a newUser object to store the incoming data.

		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//Reads the JSON data from the request body and binds it to newUser.

		sqlStatement := `
		INSERT INTO users (name, email)
		VALUES ($1, $2) RETURNING id, created_at`
		err := db.QueryRow(sqlStatement, newUser.Name, newUser.Email).Scan(&newUser.ID, &newUser.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//$1 and $2 are placeholders for the name and email values.
		//RETURNING id, created_at tells the database to return the new user's id and created_at timestamp after insertion.

		c.JSON(http.StatusCreated, newUser)
	}
}

func UpdateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var updatedUser models.User
		if err := c.ShouldBindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		sqlStatement := `
		UPDATE users
		SET name = $1, email = $2
		WHERE id = $3 RETURNING id, created_at`
		err := db.QueryRow(sqlStatement, updatedUser.Name, updatedUser.Email, id).Scan(&updatedUser.ID, &updatedUser.CreatedAt)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, updatedUser)
	}
}

func DeleteUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		sqlStatement := `DELETE FROM users WHERE id = $1`
		_, err := db.Exec(sqlStatement, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	}
}
