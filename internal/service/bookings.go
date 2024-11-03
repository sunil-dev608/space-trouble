package service

import (
	"context"
	"errors"

	"github.com/sunil-dev608/space-trouble/internal/pkg/model"
	"github.com/sunil-dev608/space-trouble/internal/repository"
)

type BookingService interface {
	CreateBooking(ctx context.Context, booking *model.BookingDB) (int64, error)
	GetAllBookings(ctx context.Context) ([]model.Booking, error)
	DeleteBooking(ctx context.Context, id int64) error
}

type bookingService struct {
	bookingRepo repository.BookingRepository
}

// NewBookingService creates a new BookingService
func NewBookingService(bookingRepo repository.BookingRepository) BookingService {
	return &bookingService{
		bookingRepo: bookingRepo,
	}
}

// CreateBooking creates a new booking in the database
// Makes 3rd party API call to check if the booking is possible
// Returns an error if the booking could not be created
// Returns an error if the booking is not valid
// Returns the ID of the newly created booking
func (s *bookingService) CreateBooking(ctx context.Context, booking *model.BookingDB) (int64, error) {

	// A requirement that can be assumed is that the launch date cannot be in the past
	// commented out for now
	// if booking.LaunchDate.Before(time.Now()) {
	// 	return -1, errors.New("launch date cannot be in the past")
	// }

	hasConflictingFlight, err := s.bookingRepo.HasConflictingFlight(ctx, booking)
	if err != nil {
		return -1, err
	}
	if hasConflictingFlight {
		return -1, errors.New("conflicting flight")
	}

	return s.bookingRepo.CreateBooking(ctx, booking)
}

// GetAllBookings returns all the bookings in the database
func (s *bookingService) GetAllBookings(ctx context.Context) ([]model.Booking, error) {
	return s.bookingRepo.GetAllBookings(ctx)
}

// DeleteBooking deletes a booking from the database by the given ID
// Returns an error if the booking is not found
// Returns gorm.ErrRecordNotFound if the booking is not found
func (s *bookingService) DeleteBooking(ctx context.Context, id int64) error {
	return s.bookingRepo.DeleteBooking(ctx, id)
}
