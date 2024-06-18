package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/null-none/gr-tours/models"
)

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}