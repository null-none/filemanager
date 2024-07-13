package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/null-none/filemanager/models"
)

// POST /files
// Create files
func CreateFile(c *gin.Context) {

	type FileInput struct {
		Folder int `json:"folder"`
	}

	var input FileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	folderStr := c.PostForm("folder")
	folder, err := strconv.Atoi(folderStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid folder value"})
		return
	}

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
		file := models.File{Name: filename, Path: fmt.Sprintf("/data/%s", filename), Folder: folder}
		models.DB.Create(&file)
	}

	c.String(http.StatusOK, "Uploaded successfully %d files.", len(files))
}
