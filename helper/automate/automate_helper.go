package automate

import (
	"auto-booking/helper"
	"auto-booking/model"
	"auto-booking/repository"
)

type Automate interface {
	BookingTennis(booking model.Booking, userBooking repository.UserBooking, sendMail helper.SendMail) error
	BookingTennisInterval(booking model.Booking, userBooking repository.UserBooking, sendMail helper.SendMail) error
}
