package app

import (
	"evara-backend/internal/config"
	"evara-backend/internal/middleware"
)

type App struct {
	server *Server
}

func New() (*App, error) {

	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	if err := middleware.InitJWKS(cfg.JWT); err != nil{
		return nil, err
	}
	
	db, err := config.NewDatabase(cfg.Database)
	
	if err != nil {
		return nil, err
	}

	providers := NewProviders(db)

	router := NewRouter(providers)

	server := NewServer(cfg.Server, router)

	return &App{
		server: server,
	}, nil
}

func (a *App) Run() error {
	return a.server.Run()
}