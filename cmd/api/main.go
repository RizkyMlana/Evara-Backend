package main

import (
	"evara-backend/internal/app"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	app, err := app.New()
	
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Run())
}