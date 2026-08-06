package middleware

import (
	"evara-backend/pkg/logger"
	"evara-backend/pkg/response"

	"github.com/gin-gonic/gin"
)


func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func ()  {
			if r := recover(); r != nil {
				logger.Log.Error(
					"panic recovered",
					"request_id", c.GetString(RequestIDKey),
					"method", c.Request.Method,
					"path", c.Request.URL.Path,
					"panic", r,
				)
				response.Internal(
					c, "internal server error",
				)

				c.Abort()
			}
		}()
		c.Next()
	}
}