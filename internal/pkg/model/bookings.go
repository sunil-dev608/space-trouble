package model

import (
	"time"
)

const (
	DateLayout       = "2006-01-02"
	BookingTableName = "bookings"
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
	ID            int64     `json:"id" gorm:"primary_key"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Gender        string    `json:"gender"`
	Birthday      time.Time `json:"birthday"`
	LaunchpadID   string    `json:"launchpad_id"`
	DestinationID string    `json:"destination_id"`
	LaunchDate    time.Time `json:"launch_date"`
}

type BookingResponse struct {
	ID int64 `json:"id"`
}

func (b *BookingDB) TableName() string {
	return "bookings"
}

// ToDB converts a Booking JSON object to a BookingDB
func (b *Booking) ToDB() (*BookingDB, error) {
	dob, err := time.Parse(DateLayout, b.Birthday)
	if err != nil {
		return nil, err
	}
	launchDate, err := time.Parse(DateLayout, b.LaunchDate)
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

func (b *Booking) Validate() bool {
	if b.FirstName == "" || b.LastName == "" || b.Gender == "" || b.Birthday == "" ||
		b.LaunchpadID == "" || b.DestinationID == "" || b.LaunchDate == "" {
		return false
	}
	return true
}
