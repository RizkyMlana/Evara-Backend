package middleware

import (
	"context"
	"evara-backend/internal/config"
	"evara-backend/pkg/response"
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
func unauthorized(c *gin.Context) {
				response.Unauthorized(c, "invalid access token")
				c.Abort()
		}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		

		if authHeader == "" {
			unauthorized(c)
			return 
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(
			tokenString,
			jwksKeyfunc,
		)
		if err != nil {
			unauthorized(c)
			return 
		}

		if !token.Valid {
			unauthorized(c)
			return 
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.Unauthorized(c, "invalid access token")
			c.Abort()
			return
		}

		userID, ok := claims["sub"].(string)
		if !ok {
			response.Unauthorized(c, "invalid access token")
			c.Abort()
			return
		}

		c.Set(UserIDKey, userID)

		c.Next()
	}
}