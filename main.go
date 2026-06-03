package main

import (
	"log"
	"os"

	"logqian-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	r := gin.Default()
	routes.SetupRoutes(r)
	log.Println("Server running on port", port)
	r.Run(":" + port)
}