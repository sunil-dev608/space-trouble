package model

import (
	"time"

	"github.com/sunil-dev608/space-trouble/config"
	"github.com/sunil-dev608/space-trouble/internal/competitors"
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

type ValidationStatus int

const (
	Valid ValidationStatus = iota
	MissingParameters
	InvalidDestination
	InvalidLaunchpad
	InvalidLaunchDate
	InvalidBirthday
)

func (v ValidationStatus) String() string {
	switch v {
	case Valid:
		return "Valid"
	case MissingParameters:
		return "MissingParameters"
	case InvalidDestination:
		return "InvalidDestination"
	case InvalidLaunchpad:
		return "InvalidLaunchpad"
	case InvalidLaunchDate:
		return "InvalidLaunchDate"
	case InvalidBirthday:
		return "InvalidBirthday"
	}
	return "Unknown"
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

func (b *Booking) Validate(cfg *config.Config) ValidationStatus {

	if b.FirstName == "" || b.LastName == "" || b.Gender == "" || b.Birthday == "" ||
		b.LaunchpadID == "" || b.DestinationID == "" || b.LaunchDate == "" {
		return MissingParameters
	}

	if _, found := cfg.Destinations[b.DestinationID]; !found {
		return InvalidDestination
	}

	if status, found := cfg.Launchpads[b.LaunchpadID]; !found || status != competitors.LaunchpadActive {
		return InvalidLaunchpad
	}

	if _, err := time.Parse(DateLayout, b.Birthday); err != nil {
		return InvalidBirthday
	}

	if _, err := time.Parse(DateLayout, b.LaunchDate); err != nil {
		return InvalidLaunchDate
	}

	return Valid
}
