package app

import (
	"evara-backend/internal/config"
	"evara-backend/internal/database"
	"evara-backend/internal/middleware"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	server *Server
	db *pgxpool.Pool
}

func New() (*App, error) {

	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	if err := middleware.InitJWKS(cfg.JWT); err != nil{
		return nil, err
	}
	
	db, err := database.NewDatabase(cfg.Database)
	
	if err != nil {
		return nil, err
	}

	providers := NewProviders(db)

	router := NewRouter(providers)

	server := NewServer(cfg.Server, router)

	return &App{
		server: server,
		db: db,
	}, nil
}

func (a *App) Run() error {
	return a.server.Run()
}
func (a *App) Close() {
	a.db.Close()
}