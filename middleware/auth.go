package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func (c * gin.Context)  {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no auth header"})
			c.Abort()
			return 
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		secret := os.Getenv("SUPABASE_JWT_SECRET")
		token, err := jwt.Parse(tokenString, func(token*jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return 
		}
		claims := token.Claims.(jwt.MapClaims)

		userID := claims["sub"]
		c.Set("user_id", userID)
		c.Next()

	}
}