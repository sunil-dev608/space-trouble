package repository

import (
	"github.com/sunil-dev608/space-trouble/internal/pkg/model"

	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(booking *model.Booking) (int64, error)
	GetAllBookings() ([]model.Booking, error)
	DeleteBooking(id uint) error
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *bookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) CreateBooking(booking *model.Booking) (int64, error) {
	dbBooking, err := booking.ToDB()
	if err != nil {
		return -1, err
	}
	result := r.db.Table(dbBooking.TableName()).Create(dbBooking)
	if result.Error != nil {
		return -1, result.Error
	}
	return dbBooking.ID, nil
}

func (r *bookingRepository) GetAllBookings() ([]model.Booking, error) {
	var bookings []model.Booking
	return bookings, r.db.Find(&bookings).Error
}

func (r *bookingRepository) DeleteBooking(id uint) error {
	return r.db.Where("id = ?", id).Delete(&model.Booking{}).Error
}
