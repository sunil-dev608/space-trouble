package server

import (
	"github.com/sunil-dev608/space-trouble/config"
	"github.com/sunil-dev608/space-trouble/internal/handlers"
	"github.com/sunil-dev608/space-trouble/internal/pkg/logger"
	"github.com/sunil-dev608/space-trouble/internal/pkg/middleware"
	"github.com/sunil-dev608/space-trouble/internal/repository"
	"github.com/sunil-dev608/space-trouble/internal/service"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func setupRoutes(e *echo.Echo, cfg *config.Config, db *gorm.DB, logger logger.Logger) {
	// Create handlers
	bookingRepo := repository.NewBookingRepository(db)
	bookingService := service.NewBookingService(cfg, bookingRepo)
	bookingsHandler := handlers.NewBookingHandler(bookingService, cfg, logger)

	// API v1 group
	v1 := e.Group("/api/v1")

	// Users routes
	users := v1.Group("/bookings")
	users.POST("/create", bookingsHandler.CreateBooking, middleware.Auth())
	users.GET("/all", bookingsHandler.GetAllBookings, middleware.Auth())
	users.DELETE("/:id", bookingsHandler.DeleteBooking, middleware.Auth())

}
