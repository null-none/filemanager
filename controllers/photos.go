package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/null-none/qr-tours/models"
)

// POST /photos
// Create photos
func CreatePhoto(c *gin.Context) {

	type PhotoInput struct {
		Tour int `json:"tour"`
	}

	var input PhotoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tour := c.PostForm("tour")

	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}
	files := form.File["files"]

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}
		photo := models.Photo{Name: filename, Path: fmt.Sprintf("/photos/%s", filename), Tour: tour}
		models.DB.Create(&photo)
	}

	c.String(http.StatusOK, "Uploaded successfully %d files.", len(files))
}
