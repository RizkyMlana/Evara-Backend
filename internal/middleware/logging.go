package middleware

import (
	"evara-backend/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func Logging() gin.HandlerFunc {
	return func (c *gin.Context)  {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		requestID := c.GetString(RequestIDKey)
		userID, _ := c.Get(UserIDKey)
		logger.Log.Info(
			"http request",
			"request_id", requestID,
			"user_id", userID,
			"method", method,
			"path", path,
			"status", status,
			"latency", latency,
			"ip", c.ClientIP(),
		)
	}
}