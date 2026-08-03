package app

import (
	"errors"
	"evara-backend/internal/config"
	"evara-backend/pkg/logger"
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
	logger.Log.Info(
		"starting HTTP server",
		"address",
		s.http.Addr,
	)

	err := s.http.ListenAndServe()
	if err != nil &&
	!errors.Is(err, http.ErrServerClosed) {
		logger.Log.Error(
			"http server stopped",
			"error",
			err,
		)

		return err
	}
	logger.Log.Info(
		"http server stopped",
	)
	return nil
}

