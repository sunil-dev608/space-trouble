package repository

import (
	"github.com/sunil-dev608/space-trouble/internal/pkg/model"

	"gorm.io/gorm"
)

// BookingRepository is the interface for the bookings repository
type BookingRepository interface {
	CreateBooking(booking *model.Booking) (int64, error)
	GetAllBookings() ([]model.Booking, error)
	DeleteBooking(id int64) error
}

type bookingRepository struct {
	db *gorm.DB
}

var bookingRepo *bookingRepository

// NewBookingRepository creates a new booking repository
func NewBookingRepository(db *gorm.DB) *bookingRepository {
	if bookingRepo != nil {
		return bookingRepo
	}
	bookingRepo = &bookingRepository{db: db}
	return bookingRepo
}

// CreateBooking creates a new booking in the database
// Returns an error if the booking could not be created
// Returns the ID of the newly created booking
func (r *bookingRepository) CreateBooking(booking *model.Booking) (int64, error) {
	dbBooking, err := booking.ToDB()
	if err != nil {
		return -1, err
	}
	result := r.db.Table(model.BookingTableName).Create(dbBooking)
	if result.Error != nil {
		return -1, result.Error
	}
	return dbBooking.ID, nil
}

// GetAllBookings returns all the bookings in the database
func (r *bookingRepository) GetAllBookings() ([]model.Booking, error) {
	var bookings []model.Booking
	return bookings, r.db.Table(model.BookingTableName).Find(&bookings).Error
}

// DeleteBooking deletes a booking from the database by the given ID
// Returns an error if the booking is not found
// Returns gorm.ErrRecordNotFound if the booking is not found
func (r *bookingRepository) DeleteBooking(id int64) error {
	result := r.db.Table(model.BookingTableName).Where("id = ?", id).Delete(&model.Booking{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
