package server

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture-example/config"
	"go-clean-architecture-example/pkg/logger"
	"os"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

// Server struct
type Server struct {
	gin    *gin.Engine
	cfg    *config.Configuration
	logger logger.Logger
}

// NewServer New Server constructor
func NewServer(cfg *config.Configuration, logger logger.Logger) *Server {
	return &Server{gin: gin.Default(), cfg: cfg, logger: logger}
}

func (s *Server) Run() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	s.logger.Info("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}
