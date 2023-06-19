package server

import (
	"context"
	"go-clean-architecture-example/config"
	"go-clean-architecture-example/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
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
	return &Server{gin: gin.New(), cfg: cfg, logger: logger}
}

func (s *Server) Run() error {
	s.gin.Use(
		gin.Recovery(),
	)
	if s.cfg.Server.Mode == "Development" {
		s.gin.Use(
			gin.Logger(),
		)
	}
	server := &http.Server{
		Addr:           s.cfg.Server.Port,
		Handler:        s.gin,
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil {
			s.logger.Fatalf("Error starting Server: ", err)
		}
	}()

	if err := s.MapHandlers(s.gin); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	s.logger.Info("Server Exited Properly")
	return server.Shutdown(ctx)
}
