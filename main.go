package main

import (
	"github.com/gin-gonic/gin"

	"github.com/null-none/qr-tours/controllers"
	"github.com/null-none/qr-tours/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/users", controllers.FindUsers)
	r.GET("/tours", controllers.FindTours)
	r.POST("/photos", controllers.CreatePhoto)

	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Static("/files", "./files")

	r.Run()
}
