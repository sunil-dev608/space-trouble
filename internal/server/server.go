package server

import (
	"context"

	"github.com/sunil-dev608/space-trouble/config"
	"github.com/sunil-dev608/space-trouble/internal/pkg/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type Server struct {
	echo   *echo.Echo
	cfg    *config.Config
	logger logger.Logger
}

func New(cfg *config.Config, db *gorm.DB, logger logger.Logger) *Server {
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Setup routes
	setupRoutes(e, cfg, db, logger)

	return &Server{
		echo:   e,
		cfg:    cfg,
		logger: logger,
	}
}

func (s *Server) Start() error {
	return s.echo.Start(s.cfg.ServerAddress)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}
