package middlewares

import (
	"gin-framework/handlers"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")

		if apiKey != os.Getenv("X_API_KEY") {
			c.JSON(401, gin.H{"error": "API key is invalid or missing"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func JWTAUth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := c.GetHeader("Authorization")
		claims := &handlers.Claims{}

		tkn, err := jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		if tkn == nil || !tkn.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
