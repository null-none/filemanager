package main

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/null-none/qr-tours/models"
	"github.com/null-none/qr-tours/controllers"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Static("/", "./")

	r.GET("/users", controllers.FindUsers)
	r.GET("/tours", controllers.FindTours)
	r.POST("/photos", controllers.CreatePhoto)

	r.Run()
}