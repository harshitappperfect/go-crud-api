package handlers

import (
	"database/sql"
	"go-crud-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary Get all users
// @Description Get a list of all users from the database
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {string} string "error"
// @Failure 404 {string} string "error"
// @Router /users [get]
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

// GetUser godoc
// @Summary Get a specific user
// @Summary Get a users
// @Description Get a user from the database based on id
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 201 {object} models.User
// @Failure 400 {string} string "error"
// @Failure 500 {string} string "error"
// @Router /users/{id} [get]
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

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with name and email
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 201 {object} models.User
// @Failure 400 {string} string "error"
// @Failure 500 {string} string "error"
// @Router /users [post]
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

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user's name and email
// @Tags users
// @Accept json
// @produce json
// @Param id path string true "User ID"
// @Param user body models.User true "Updated User Data"
// @Success 201 {object} models.User
// @Failure 400 {string} string "error"
// @Failure 500 {string} string "error"
// @Router /users/{id} [put]
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

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Succeess 200 {string} string {"message" : "Successfully deleted user"}
// @Failure 500 {string} string "error"
// @Failure 404 {string} string "error"
// @Router /users/{id} [delete]
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
