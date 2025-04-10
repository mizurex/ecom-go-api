package handler

import (
	"database/sql"
	"ecom/database"
	"ecom/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

func SetDB(db *sql.DB) {
	DB = db
}

func AddUserHandler(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := database.InsertUser(DB, newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user added successfully!"})

}
