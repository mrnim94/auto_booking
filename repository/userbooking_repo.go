package repository

import (
	"auto-booking/model"
	"context"
)

type UserBooking interface {
	SelectAllUsersBooking(context context.Context, amenities string) ([]model.Booking, error)
	CheckUsersBooked(context context.Context, check model.Booking) (model.Booking, error)
	UpdateUsersBooking(context context.Context, user model.Booking) (model.Booking, error)
}