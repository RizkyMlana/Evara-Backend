package routes

import (
	"logqian-backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	// api.Use(middleware.AuthMiddleware())

	api.POST("/transactions", handlers.CreateTransaction)
	api.GET("/transactions", handlers.GetTransactions)
}