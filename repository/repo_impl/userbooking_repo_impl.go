package repo_impl

import (
	"auto-booking/banana"
	"auto-booking/db"
	"auto-booking/log"
	"auto-booking/model"
	"auto-booking/repository"
	"context"
	"database/sql"
	"time"
)

type UserBookingRepoImpl struct {
	sql *db.Sql
}

func NewUserBooking(sql *db.Sql) repository.UserBooking {
	return &UserBookingRepoImpl{
		sql: sql,
	}
}

func (u *UserBookingRepoImpl) SelectAllUsersBooking(context context.Context, amenities string) ([]model.Booking, error) {
	users := []model.Booking{}
	var dayBooking int
	hourCurrent := time.Now().Hour()

	if hourCurrent == 23 {
		dayBooking = time.Now().AddDate(0, 0, +8).Day()

	} else {
		dayBooking = time.Now().AddDate(0, 0, +7).Day()
	}

	err := u.sql.Db.SelectContext(context, &users,
		`SELECT user_id, username, password, date_booking, time_booking, time_booking_backup, amenities, amenity, homes, home_id
				FROM users_booking 
				WHERE amenities=$1 AND date_booking=$2`, amenities, dayBooking)

	if err != nil {
		if err == sql.ErrNoRows {
			return users, banana.UserBookingNotFound
		}
		log.Error(err.Error())
		return users, err
	}
	return users, nil
}
func (u *UserBookingRepoImpl) CheckUsersBooked(context context.Context, check model.Booking) (model.Booking, error)  {
	user := model.Booking{}
	err := u.sql.Db.GetContext(context, &user, "SELECT * FROM users_booking " +
		"WHERE username=$1 AND amenity=$2 AND date_booking=$3", check.Username, check.Amenity, check.DateBooking)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, banana.UserBookingNotFound
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}

func (u *UserBookingRepoImpl) UpdateUsersBooking(context context.Context, user model.Booking) (model.Booking, error) {
	sqlStatement := `
		UPDATE users_booking
		SET 
			date_booking  = (CASE WHEN LENGTH(:date_booking) = 0 THEN date_booking ELSE :date_booking END),
			created_booking  = (CASE WHEN LENGTH(:created_booking) = 0 THEN created_booking ELSE :created_booking END),
			completed_booking = (CASE WHEN LENGTH(:completed_booking) = 0 THEN completed_booking ELSE :completed_booking END)
		WHERE user_id = :user_id
	`

	result, err := u.sql.Db.NamedExecContext(context, sqlStatement, user)
	if err != nil {
		log.Error(err.Error())
		return user, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return user, banana.UserNotUpdated
	}
	if count == 0 {
		return user, banana.UserBookingNotFound
	}

	return user, nil

}