package main

import (
	"log"
	"os"

	"evara-backend/config"
	"evara-backend/middleware"
	"evara-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := middleware.InitJWKS(); err != nil {
		log.Fatal(err)
	}

	config.ConnectDB()
	port := os.Getenv("PORT")
	r := gin.Default()
	routes.SetupRoutes(r)
	log.Println("Server running on port", port)
	r.Run(":" + port)
}