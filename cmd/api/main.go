// @title Evara API
// @version 1.0
// @description Family Expense Tracker API
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"evara-backend/internal/app"
	"evara-backend/pkg/logger"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	logger.Init()
	app, err := app.New()
	
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Run())
}