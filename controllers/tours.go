package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/null-none/gr-tours/models"
)

// GET /tours
// Get all tours
func FindTours(c *gin.Context) {
	var tours []models.Tour
	models.DB.Find(&tours)

	c.JSON(http.StatusOK, gin.H{"data": users})
}