package main

import (
	"github.com/gin-gonic/gin"

	"github.com/null-none/filemanager/controllers"
	"github.com/null-none/filemanager/models"
)

func main() {
	router:= gin.Default()

	models.ConnectDatabase()
	router.POST("/login", controllers.Login)

	router.GET("/folders", middleware.authMiddleware(), controllers.FindFolders)
	router.GET("/folders/:id", middleware.authMiddleware(), controllers.FindFolder)
	
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/data", "./data")

	router.Run()
}