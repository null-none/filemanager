package controllers

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/gin-gonic/gin"
	"github.com/null-none/filemanager/models"
	"github.com/null-none/filemanager/middlewares"
)


// POST /login
// Login
func Login(c *gin.Context) {

	type Validate struct {
		Phone     string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
		Password  string `form:"password" json:"password" xml:"password" binding:"required"`
	}

	var user models.User
  
	var json Validate
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("phone = ?", json.Phone).First(&user).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}

    tokenString, err := middlewares.CreateToken(user.Phone)
    if err != nil {
       w.WriteHeader(http.StatusInternalServerError)
       fmt.Errorf(err)
     }

	c.JSON(http.StatusOK, gin.H{"data": user, "token": tokenString})
}