package service

import (
	"github.com/sunil-dev608/space-trouble/internal/pkg/model"
	"github.com/sunil-dev608/space-trouble/internal/repository"
)

type BookingService interface {
	CreateBooking(booking *model.Booking) (int64, error)
	GetAllBookings() ([]model.Booking, error)
	DeleteBooking(id int64) error
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
func (s *bookingService) CreateBooking(booking *model.Booking) (int64, error) {
	return s.bookingRepo.CreateBooking(booking)
}

// GetAllBookings returns all the bookings in the database
func (s *bookingService) GetAllBookings() ([]model.Booking, error) {
	return s.bookingRepo.GetAllBookings()
}

// DeleteBooking deletes a booking from the database by the given ID
// Returns an error if the booking is not found
// Returns gorm.ErrRecordNotFound if the booking is not found
func (s *bookingService) DeleteBooking(id int64) error {
	return s.bookingRepo.DeleteBooking(id)
}
