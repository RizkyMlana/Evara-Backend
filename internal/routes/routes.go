package routes

import (
	"evara-backend/internal/family"
	"evara-backend/internal/middleware"
	"evara-backend/internal/transaction"

	"github.com/gin-gonic/gin"
)


func SetupRoutes(
	r *gin.Engine,
	transactionHandler *transaction.Handler,
	familyHandler *family.Handler,
	) {

		api := r.Group("/api")
		api.Use(
			middleware.AuthMiddleware(),
		)

		api.POST(
			"/transactions",
			transactionHandler.CreateTransaction,
		)
}