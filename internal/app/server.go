package app

import (
	"evara-backend/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	http *http.Server
}
func NewServer(
	cfg config.ServerConfig,
	router *gin.Engine,
) *Server {
	server := &http.Server{
		Addr: ":" + cfg.Port,
		Handler: router,
	}

	return &Server{
		http: server,
	}
}

func (s *Server) Run() error {
	return s.http.ListenAndServe()
}

