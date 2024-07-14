package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/null-none/filemanager/models"
)

// POST /login
// Login
func Login(c *gin.Context) {


	type Validate struct {
		Login     string `form:"login" json:"login" xml:"login"  binding:"required"`
		Password  string `form:"password" json:"password" xml:"password" binding:"required"`
	}

	var user models.User

	var json Validate
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter your username and password."})
		return
	}

	if err := models.DB.Where("login = ?", json.Login).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found."})
		return
	}

	token, err := createToken(user.Login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization error, please try again later."})
	}

	c.JSON(http.StatusOK, gin.H{"data": user, "token": token})

}
