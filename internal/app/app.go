package app

import (
	"net/http"
	"os"

	"evara-backend/internal/config"
	"evara-backend/internal/family"
	"evara-backend/internal/routes"
	"evara-backend/internal/transaction"

	"github.com/gin-gonic/gin"
)

type App struct{
	router *gin.Engine
	server *http.Server
}

func New() (*App, error) {
	db, err := config.NewDatabase()
	if err != nil {
		return nil, err
	}

	// transaction
	transactionRepository := transaction.NewRepository(db)
	transactionService := transaction.NewService(transactionRepository)
	transactionHandler := transaction.NewHandler(transactionService)

	// family
	familyRepository := family.NewRepository(db)
	familyService := family.NewService(familyRepository)
	familyHandler := family.NewHandler(familyService)

	router := gin.Default()


	router.SetTrustedProxies(
		[]string{"192.168.1.2"},
	)

	routes.SetupRoutes(
		router,
		transactionHandler,
		familyHandler,
	)

	server := &http.Server{
		Addr: ":" + os.Getenv("PORT"),
		Handler: router,
	}

	return &App{
		router: router,
		server: server,
	}, nil
}

func (a *App) Run() error {
	return a.server.ListenAndServe()
}