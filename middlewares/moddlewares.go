package middleware


import (
	"time"
    "net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/null-none/filemanager/utils"
	"github.com/golang-jwt/jwt/v5"
)

func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            return
        }

		err := token.verifyToken(tokenString)
		
		if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}

        c.Next()
    }
}