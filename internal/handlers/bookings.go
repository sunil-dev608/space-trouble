package handlers

import (
	"errors"
	"net/http"
	"strconv"

	// "your-project/pkg/logger"
	// "your-project/pkg/response"

	"github.com/sunil-dev608/space-trouble/internal/pkg/logger"
	"github.com/sunil-dev608/space-trouble/internal/pkg/model"
	"github.com/sunil-dev608/space-trouble/internal/pkg/response"
	"github.com/sunil-dev608/space-trouble/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/sunil-dev608/space-trouble/config"
)

// BookingHandler is the API handler for bookings
type BookingHandler struct {
	bookingService service.BookingService
	cfg            *config.Config
	logger         logger.Logger
}

// NewBookingHandler returns a new instance of BookingHandler
func NewBookingHandler(bookingService service.BookingService, cfg *config.Config, logger logger.Logger) *BookingHandler {
	return &BookingHandler{
		bookingService: bookingService,
		cfg:            cfg,
		logger:         logger,
	}
}

// CreateBooking creates a new booking
func (h *BookingHandler) CreateBooking(c echo.Context) error {
	var req model.Booking
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request", err)
	}

	if !req.Validate() {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request", errors.New("all fields are required"))
	}

	if id, err := h.bookingService.CreateBooking(&req); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, "failed to create booking", err)
	} else {
		resp := model.BookingResponse{ID: id}

		return response.SuccessResponse(c, http.StatusCreated, "user created", &resp)
	}

}

// GetAllBookings returns all the bookings
func (h *BookingHandler) GetAllBookings(c echo.Context) error {
	bookings, err := h.bookingService.GetAllBookings()
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, "failed to get bookings", err)
	}
	return response.SuccessResponse(c, http.StatusOK, "bookings retrieved", bookings)
}

// DeleteBooking deletes a booking
func (h *BookingHandler) DeleteBooking(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request", errors.New("id is required"))
	}

	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request", err)
	}

	if err := h.bookingService.DeleteBooking(ID); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, "failed to delete booking", err)
	} else {
		return response.SuccessResponse(c, http.StatusOK, "booking deleted", nil)
	}
}
