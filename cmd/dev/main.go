package main

import (
	"auto-booking/db"
	"auto-booking/handler"
	"auto-booking/helper"
	"auto-booking/helper/automate/automate_impl"
	"auto-booking/log"
	"auto-booking/repository/repo_impl"
	"auto-booking/router"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
	"os"
)

func init() {
	os.Setenv("APP_NAME", "backend_codedeploy_dev")
	log.InitLogger(false)
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
}

func main()  {

	sql := &db.Sql{
		Host:     "192.168.101.5", //localhost,host.docker.internal
		Port:     5432,
		UserName: "nim",
		Password: "nim123",
		DbName:   "autobooking",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()

	seConNim := &automate_impl.Selenium{
		Browser:       "chrome",
		//ConnectServer: "http://192.168.101.35:4444/wd/hub",
		ConnectServer: "http://192.168.101.9:4444/wd/hub",
	}

	conMail := &helper.AccountMail{
		Host:     "smtp.gmail.com",
		Port:     587,
		Username: "autobooking.nim@gmail.com",
		Password: "123456789nim",
	}

	autoBookingHandler := handler.AutoBookingHandler{
		Automate: automate_impl.NewSelenium(seConNim),
		UserBooking: repo_impl.NewUserBooking(sql),
		SendMail: helper.NewSendMail(conMail),
	}

	api := router.API{
		Echo:           e,
		AutoBookingHandler: autoBookingHandler,
	}
	api.SetupRouter()

	//autoBookingHandler.HandlerScheduleBookingTennis()

	c := cron.New(cron.WithSeconds())

	c.AddFunc("CRON_TZ=Asia/Ho_Chi_Minh 30 59 23 * * *", func() {
		for i := 1; i <= 2; i++ {
			go autoBookingHandler.HandlerScheduleBookingTennis()
		}
	})


	//Khu vực Booking lập lại
	var interval30s cron.EntryID

	c.AddFunc("CRON_TZ=Asia/Ho_Chi_Minh 00 00 00 * * *", func() {
		interval30s, _ = c.AddFunc("@every 0h0m30s", func() { autoBookingHandler.HandlerIntervalBookingTennis() })
	})

	c.AddFunc("CRON_TZ=Asia/Ho_Chi_Minh 00 05 09 * * *", func() { c.Remove(interval30s) })

	c.Start()


	e.Logger.Fatal(e.Start(":6969"))
}