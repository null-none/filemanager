package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/null-none/filemanager/models"
)

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GET /users/:id
// Find a user
func FindUser(c *gin.Context) {  // Get model if exist
	var user models.User
  
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	c.JSON(http.StatusOK, gin.H{"data": user})
  }