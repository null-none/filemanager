package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/null-none/filemanager/models"
)

// GET /folders
// Get all folders
func FindFolders(c *gin.Context) {
	var folders []models.Folder
	models.DB.Find(&folders)

	c.JSON(http.StatusOK, gin.H{"data": folders})
}

// GET /folders/:id
// Find a folder
func FindFolder(c *gin.Context) { // Get model if exist
	var folder models.Folder

	if err := models.DB.Where("id = ?", c.Param("id")).First(&folder).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Folder not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": folder})
}
