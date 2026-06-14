package routes

import (
	"evara-backend/handlers"
	"evara-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())

	api.GET("/me", handlers.Me)

	api.POST("/families", handlers.CreateFamily)
	api.GET("/families/me", handlers.GetMyFamily)
	api.GET("/families/:id/members", handlers.GetFamilyMember)
	api.POST("/families/:id/invite", handlers.InviteMember)
	api.GET("/invitations", handlers.GetInvitations)
	api.POST("/invitations/:id/accept", handlers.AcceptInvitations)
	api.POST("/invitations/:id/reject", handlers.RejectInvitation)

	api.POST("/transaction", handlers.CreateTransaction)
	api.GET("/transaction", handlers.GetTransactions)
}