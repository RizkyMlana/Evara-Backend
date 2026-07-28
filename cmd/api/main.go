package main

import (
	"log"
	"os"

	"evara-backend/internal/config"
	"evara-backend/internal/middleware"
	"evara-backend/internal/routes"
	"evara-backend/internal/transaction"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := middleware.InitJWKS(); err !=nil {
		log.Fatal(err)
	}

	config.ConnectDB()

	transactionRepository := transaction.NewRepository()

	transactionService := transaction.NewService(
		transactionRepository,
	)

	transactionHandler := transaction.NewHandler(
		transactionService,
	)

	port := os.Getenv("PORT")

	r:= gin.Default()

	r.SetTrustedProxies(
		[]string{"192.168.1.2"},
	)

	routes.SetupRoutes(
		r,
		transactionHandler,
	)

	log.Println("Server running on port", port)

	r.Run(":"+port)
}