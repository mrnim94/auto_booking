package handler

import (
	"auto-booking/banana"
	"auto-booking/helper"
	"auto-booking/helper/automate"
	"auto-booking/log"
	"auto-booking/model"
	"auto-booking/repository"
	"context"
	"github.com/labstack/echo/v4"
	"time"
)

type AutoBookingHandler struct {
	Echo        *echo.Context
	Automate automate.Automate
	UserBooking repository.UserBooking
	SendMail helper.SendMail
}

func (a *AutoBookingHandler) HandlerScheduleBookingTennis() error {
	log.Printf("Khởi động auto Schedule Booking")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	users, err := a.UserBooking.SelectAllUsersBooking(ctx, "TENNIS NEW")
	if err != nil {
		if err == banana.UserBookingNotFound {
			log.Error(err.Error())
			return banana.UserBookingNotFound
		}
		log.Error(err.Error())
		return err
	}
	if len(users) == 0 {
		log.Info("Hệ thống không lấy được user nào")
	}else {
		for i, user := range users {
			updateBooking := model.Booking{
				UserId:           user.UserId,
				CreatedBooking:   time.Now(),
			}
			_, err = a.UserBooking.UpdateUsersBooking(ctx, updateBooking)
			if err != nil {
				log.Error(err.Error())
				return err
			}

			user.CreatedBooking = time.Now()
			go a.Automate.BookingTennis(user, a.UserBooking, a.SendMail)
			log.Printf("Thông tin users lấy được", users[i])
		}
	}

	return nil
}


func (a *AutoBookingHandler) HandlerIntervalBookingTennis() error {
	log.Printf("Khởi động auto Interval Booking")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	users, err := a.UserBooking.SelectAllUsersBooking(ctx, "TENNIS NEW")

	if err != nil {
		if err == banana.UserBookingNotFound {
			log.Error(err.Error())
			return banana.UserBookingNotFound
		}
		log.Error(err.Error())
		return err
	}
	if len(users) == 0 {
		log.Info("Hệ thống không lấy được user nào")
	}else {
		for i, user := range users {
			updateBooking := model.Booking{
				UserId:           user.UserId,
				CreatedBooking:   time.Now(),
			}
			_, err = a.UserBooking.UpdateUsersBooking(ctx, updateBooking)
			if err != nil {
				log.Error(err.Error())
				return err
			}

			user.CreatedBooking = time.Now()
			go a.Automate.BookingTennisInterval(user, a.UserBooking, a.SendMail)
			log.Printf("Thông tin users lấy được", users[i])
		}
	}

	return nil
}