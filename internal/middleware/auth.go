package middleware

import (
	"context"
	"evara-backend/internal/config"
	"net/http"
	"strings"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwksKeyfunc jwt.Keyfunc

func InitJWKS(cfg config.JWTConfig) error {
	jwks, err := keyfunc.NewDefaultCtx(
		context.Background(),
		[]string{
			cfg.URL,
		},
	)

	if err != nil {
		return err
	}

	jwksKeyfunc = jwks.Keyfunc


	return nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "no auth header",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(
			tokenString,
			jwksKeyfunc,
		)

		if err != nil {

			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token not valid",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid claims",
			})
			c.Abort()
			return
		}



		userID, ok := claims["sub"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing user id",
			})
			c.Abort()
			return
		}

		c.Set("user_id", userID)

		c.Next()
	}
}