package service

import (
	"github.com/sunil-dev608/space-trouble/internal/pkg/model"
	"github.com/sunil-dev608/space-trouble/internal/pkg/repository"
)

type BookingService interface {
	CreateBooking(booking *model.Booking) (int64, error)
	GetAllBookings() ([]model.Booking, error)
	DeleteBooking(id uint) error
}

type bookingService struct {
	bookingRepo repository.BookingRepository
}

func NewBookingService(bookingRepo repository.BookingRepository) BookingService {
	return &bookingService{
		bookingRepo: bookingRepo,
	}
}

func (s *bookingService) CreateBooking(booking *model.Booking) (int64, error) {
	return s.bookingRepo.CreateBooking(booking)
}

func (s *bookingService) GetAllBookings() ([]model.Booking, error) {
	return s.bookingRepo.GetAllBookings()
}

func (s *bookingService) DeleteBooking(id uint) error {
	return s.bookingRepo.DeleteBooking(id)
}
