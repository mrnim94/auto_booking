package model

import "time"

type Booking struct {
	UserId    string    `json:"-" db:"user_id, omitempty"`
	Username string `json:"userName,omitempty" db:"username, omitempty"`
	Password string `json:"password,omitempty" db:"password, omitempty"`
	DateBooking string `json:"dateBooking,omitempty" db:"date_booking, omitempty"`
	TimeBooking string `json:"timeBooking,omitempty" db:"time_booking, omitempty"`
	TimeBookingBackup string `json:"timeBookingBackup,omitempty" db:"time_booking_backup, omitempty"`
	Amenities string `json:"amenities,omitempty" db:"amenities, omitempty"`
	Amenity string `json:"amenity,omitempty" db:"amenity, omitempty"`
	Homes string `json:"homes,omitempty" db:"homes, omitempty"`
	HomeID string `json:"homeID,omitempty" db:"home_id, omitempty"`
	CreatedBooking time.Time `json:"createdBooking,omitempty" db:"created_booking, omitempty"`
	CompletedBooking time.Time `json:"completedBooking,omitempty" db:"completed_booking, omitempty"`
}

type StatusTimesBooking struct {
	StatusTimeBooking string `json:"statusTimeBooking,omitempty"`
}
