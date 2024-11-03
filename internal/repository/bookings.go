package repository

import (
	"context"

	"github.com/sunil-dev608/space-trouble/internal/pkg/model"

	"gorm.io/gorm"
)

// BookingRepository is the interface for the bookings repository
type BookingRepository interface {
	CreateBooking(ctx context.Context, booking *model.BookingDB) (int64, error)
	GetAllBookings(ctx context.Context) ([]model.Booking, error)
	DeleteBooking(ctx context.Context, id int64) error

	HasConflictingFlight(ctx context.Context, booking *model.BookingDB) (bool, error)
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
func (r *bookingRepository) CreateBooking(ctx context.Context, booking *model.BookingDB) (int64, error) {
	result := r.db.WithContext(ctx).Table(model.BookingTableName).Create(booking)
	if result.Error != nil {
		return -1, result.Error
	}
	return booking.ID, nil
}

// GetAllBookings returns all the bookings in the database
func (r *bookingRepository) GetAllBookings(ctx context.Context) ([]model.Booking, error) {
	var bookings []model.Booking
	return bookings, r.db.WithContext(ctx).Table(model.BookingTableName).Find(&bookings).Error
}

// DeleteBooking deletes a booking from the database by the given ID
// Returns an error if the booking is not found
// Returns gorm.ErrRecordNotFound if the booking is not found
func (r *bookingRepository) DeleteBooking(ctx context.Context, id int64) error {
	result := r.db.WithContext(ctx).Table(model.BookingTableName).Where("id = ?", id).Delete(&model.Booking{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// HasConflictingFlight checks if there is a conflicting flight
// Returns true if there is a conflicting flight
// Returns an error if the query failed
// A flight is conflicting if there is at least one flight from the same launchpad on the day before and the day after the launch date
// Ex: If the launch date is 2022-01-01 and the launchpad is 5, there is a flight on 2022-01-02 from launchpad 5 to Destination 1 then this should return true
// Ex: If the launch date is 2022-01-01 and the launchpad is 5, there is no flight on 2021-12-31 from launchpad 5 to Destination 1 then this should return true
// Ex: If the launch date is 2022-01-01 and the launchpad is 5, there is a flight on 2022-01-01 launchpad 5 to Destination 2 then this should return false
//
// Ex: If the launch date is 2022-01-01 and the launchpad is 5,
//
//	there is no flight on 2021-12-31 and 2022-01-02 from launchpad 5 to Destination 1
//	and there is no flight on 2022-01-01 from launchpad 5 to any Destination
//	then this should return false
func (r *bookingRepository) HasConflictingFlight(ctx context.Context, booking *model.BookingDB) (bool, error) {
	var count int64
	toDate := booking.LaunchDate.AddDate(0, 0, 1)
	fromDate := booking.LaunchDate.AddDate(0, 0, -1)
	result := r.db.WithContext(ctx).Table(model.BookingTableName).
		Where(`(LAUNCHPAD_ID = ? AND DESTINATION_ID = ? AND (LAUNCH_DATE = ? OR LAUNCH_DATE = ?)) OR
		 		(LAUNCHPAD_ID = ? AND DESTINATION_ID != ? AND LAUNCH_DATE = ?)`,
			booking.LaunchpadID, booking.DestinationID, fromDate, toDate,
			booking.LaunchpadID, booking.DestinationID, booking.LaunchDate).Count(&count)
	return count > 0, result.Error
}
