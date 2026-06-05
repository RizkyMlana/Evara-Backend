package routes

import (
	"logqian-backend/handlers"
	"logqian-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())

	api.POST("/transactions", handlers.CreateTransaction)
	api.GET("/transactions", handlers.GetTransactions)

	api.GET("/me", handlers.Me)

	api.POST("/families", handlers.CreateFamily)
	api.GET("/families/me", handlers.GetMyFamily)
	api.GET("/families/:id/members", handlers.GetFamilyMember)
	api.POST("/families/:id/invite")
	api.GET("/invitations", handlers.GetInvitations)
}