package main

import (
	"github.com/gin-gonic/gin"

	"github.com/null-none/filemanager/controllers"
	"github.com/null-none/filemanager/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/users", controllers.FindUsers)
	r.GET("/users/:id", controllers.FindUser)
	r.GET("/tours", controllers.FindTours)
	r.POST("/files", controllers.CreateFile)

	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Static("/data", "./data")

	r.Run()
}
