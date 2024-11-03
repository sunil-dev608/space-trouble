package model

import (
	"time"

	"github.com/sunil-dev608/space-trouble/internal/pkg/constants"
)

type Booking struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Gender        string `json:"gender"`
	Birthday      string `json:"birthday"`
	LaunchpadID   string `json:"launchpad_id"`
	DestinationID string `json:"destination_id"`
	LaunchDate    string `json:"launch_date"`
}

type BookingDB struct {
	// gorm.Model
	ID            int64     `json:"id" gorm:"primary_key"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Gender        string    `json:"gender"`
	Birthday      time.Time `json:"birthday"`
	LaunchpadID   string    `json:"launchpad_id"`
	DestinationID string    `json:"destination_id"`
	LaunchDate    time.Time `json:"launch_date"`
}

func (b *BookingDB) TableName() string {
	return "bookings"
}

func (b *Booking) ToDB() (*BookingDB, error) {
	dob, err := time.Parse(constants.DateLayout, b.Birthday)
	if err != nil {
		return nil, err
	}
	launchDate, err := time.Parse(constants.DateLayout, b.LaunchDate)
	if err != nil {
		return nil, err
	}
	return &BookingDB{
		FirstName:     b.FirstName,
		LastName:      b.LastName,
		Gender:        b.Gender,
		Birthday:      dob,
		LaunchpadID:   b.LaunchpadID,
		DestinationID: b.DestinationID,
		LaunchDate:    launchDate,
	}, nil
}
