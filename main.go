package main

import (
	"github.com/gin-gonic/gin"
    "net/http"
	"fmt"
	
	"github.com/golang-jwt/jwt/v5"
	"github.com/null-none/filemanager/controllers"
	"github.com/null-none/filemanager/models"
)

func main() {
	router:= gin.Default()

	models.ConnectDatabase()
	router.POST("/login", controllers.Login)

	router.GET("/folders", authMiddleware(), controllers.FindFolders)
	router.GET("/folders/:id", authMiddleware(), controllers.FindFolder)
	
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/data", "./data")

	router.Run()
}

func verifyToken(tokenString string) error {
	var secretKey = []byte("secret-key")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}


func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            return
        }

		err := verifyToken(tokenString)
		
		if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}

        c.Next()
    }
}