package automate_impl

import (
	"auto-booking/helper"
	"auto-booking/log"
	"auto-booking/model"
	"auto-booking/repository"
	"context"
	"fmt"
	"github.com/tebeka/selenium"
	"strconv"
	"time"
	"github.com/robfig/cron/v3"
)


func (s *Selenium) BookingTennis(booking model.Booking, userBooking repository.UserBooking, sendMail helper.SendMail) error {

	conctRemote, err := s.getRemote()
	if err != nil {
		log.Error(err.Error()+">>>"+booking.Username)
		return err
	}
	//defer conctRemote.Quit()

	err = conctRemote.Get("https://vn.propertycube.asia/account")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	conctRemote.SetImplicitWaitTimeout(time.Second * 60)

	if title, err := conctRemote.Title(); err == nil {
		fmt.Printf("Page title: %s\n", title)
	} else {
		log.Error(err.Error())
		return err
	}

	conctRemote.ResizeWindow("note", 1920, 1080)

	elemUserNameOrEmailAddress, err := conctRemote.FindElement(selenium.ByName, "userNameOrEmailAddress")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemUserNameOrEmailAddress.SendKeys(booking.Username)
	log.Printf(booking.Username + "Đã điền được User")

	elemPassword, err := conctRemote.FindElement(selenium.ByName, "password")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemPassword.SendKeys(booking.Password)
	log.Printf(booking.Username + " Đã điền được Password")

	elemBtLogin, err := conctRemote.FindElement(selenium.ByCSSSelector, "button[type=\"submit\"]")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemBtLogin.Click()
	log.Printf(booking.Username + " Đã đã nhấn login")

	elemTenancy, err := conctRemote.FindElement(selenium.ByCSSSelector, "[title='MASTERI']")
	if err != nil {
		log.Info(err.Error())
		log.Info(booking.Username + "Chỉ có Thuê 1 căn")
	}else {
		elemTenancy.Click()
		log.Printf(booking.Username + " Đã nhấn vào Tenancy MASTERI")
	}


	elemUserProfile, err := conctRemote.FindElement(selenium.ByCSSSelector, "img.rounded.ng-star-inserted")
	if err != nil {
		log.Error(err.Error())
		log.Error(booking.Username + " login thất bại")
		return err
	}
	elemUserProfile.Click()
	log.Printf(booking.Username + " Đã nhấn vào Picture Profile")

	elemCheckUserProfile, err := conctRemote.FindElement(selenium.ByCSSSelector, ".text-primary.hide-mobile")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	name, _ := elemCheckUserProfile.Text()
	log.Printf(booking.Username + " so sánh với "+ name)

	elemCheckLanguage, err := conctRemote.FindElement(selenium.ByCSSSelector, ".ml-2.hide-mobile")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	language, err := elemCheckLanguage.Text()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Printf(booking.Username + " đang ở ngôn ngữ "+ language)

	if language != "English" {
		elemChangeLanguage, err := conctRemote.FindElement(selenium.ByCSSSelector, ".flag-icon.flag-icon-vn")
		if err != nil {
			log.Error(err.Error())
			return err
		}
		elemChangeLanguage.Click()
		log.Printf(booking.Username + " chọn Change language")

		elemChooseEnglish, err := conctRemote.FindElement(selenium.ByCSSSelector, ".flag-icon.flag-icon-gb")
		if err != nil {
			log.Error(err.Error())
			return err
		}
		elemChooseEnglish.Click()
		log.Printf(booking.Username + " chọn language English")
	}

	if booking.Homes == "yes" {
		elemHomes, err :=  conctRemote.FindElement(selenium.ByCSSSelector, "span.ng-star-inserted")
		if err != nil {
			log.Error(err.Error())
			return err
		}
		elemHomes.Click()
		log.Printf(booking.Username + " chọn Homes")

		HomeId, err :=  conctRemote.FindElement(selenium.ByXPATH, "//li[contains(text(),'"+ booking.HomeID+"')]")
		if err != nil {
			log.Error(err.Error())
			return err
		}
		HomeId.Click()
		log.Printf(booking.Username + " chọn Homes")
	}


	time.Sleep(time.Second * 5)
	conctRemote.SetImplicitWaitTimeout(time.Second * 60)

	elemBooking, err := conctRemote.FindElement(selenium.ByCSSSelector, "#menuItemFacilityManagement " +
																"> li.dropdown.ng-star-inserted > a > span.menu-name")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemBooking.Click()
	log.Printf(booking.Username + " Đã nhấn vào Booking cha")

	elemBookingChild, err := conctRemote.FindElement(selenium.ByCSSSelector, "#menuItemFacilityManagement " +
										"> li.dropdown.ng-star-inserted > ul > li:nth-child(1) > a > span.menu-name")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemBookingChild.Click()
	log.Printf(booking.Username + " Đã nhấn vào Booking con")

	var amenities string
	var amenity string
	switch booking.Amenities {
	case "TENNIS NEW":
		amenities = "a[href='#amenityGroup14']"
		if booking.Amenity == "Tennis 1 New" {
			amenity = "label[for=\"amenities62\"]"
		}else{
			amenity = "label[for=\"amenities64\"]"
		}
	}
	
	elemAmenities, err := conctRemote.FindElement(selenium.ByCSSSelector, amenities)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemAmenities.Click()
	log.Printf(booking.Username + " Đã nhấn vào Amenities " + booking.Amenities)

	elemAmenity, err := conctRemote.FindElement(selenium.ByCSSSelector, amenity)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemAmenity.Click()
	log.Printf(booking.Username + " Đã đã nhấn vào Amenity " + booking.Amenity)

	conctRemote.SetImplicitWaitTimeout(time.Second * 10)

	// chon ngay
	c := cron.New()
	c.AddFunc("CRON_TZ=Asia/Ho_Chi_Minh 00 00 * * *", func() {
		after00h(conctRemote , booking , userBooking , sendMail)
	})
	c.Start()

	return nil
}

func after00h(conctRemote selenium.WebDriver, booking model.Booking, userBooking repository.UserBooking, sendMail helper.SendMail) error {
	defer conctRemote.Quit()
	log.Printf(booking.Username + "Bước vào tranh gianh Booking")
	ctx := context.Background()
	elemDateBooking, err := conctRemote.FindElement(selenium.ByXPATH,
		"//span[text()='"+booking.DateBooking+"']")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemDateBooking.Click()
	log.Printf(booking.Username + " Đã chọn ngày " + booking.DateBooking)

	//elemGetTimesUnavailable, err := conctRemote.FindElements(selenium.ByCSSSelector, "td.fc-event-container a.unavailable")
	//if err != nil {
	//	log.Error(err.Error())
	//	return err
	//}
	//
	//var successTimeBooking string
	//if len(elemGetTimesUnavailable) > 0 {
	//	for _, timeUnavailable := range elemGetTimesUnavailable {
	//		timeU, err := timeUnavailable.Text()
	//		if err != nil {
	//			log.Error(err.Error())
	//			return err
	//		}
	//
	//		if timeU == booking.TimeBooking {
	//			log.Error(booking.Username + " không thể chọn " + booking.TimeBooking)
	//			//check xem là đã có session nào book thành công chưa?
	//			checkUserBooked := model.Booking{
	//				Username:          booking.Username,
	//				DateBooking:       booking.DateBooking,
	//				Amenity:           booking.Amenity,
	//			}
	//
	//			_, err := userBooking.CheckUsersBooked(ctx, checkUserBooked)
	//			if err != nil {
	//				log.Printf(booking.Username + " đã booking thành công " + booking.TimeBooking)
	//				log.Error(err.Error())
	//				return err
	//			}else {
	//				log.Printf(booking.Username + " Đã chưa booking thành công " + booking.TimeBooking +
	//					" Chuyển sang " + booking.TimeBookingBackup)
	//				successTimeBooking = booking.TimeBookingBackup
	//				break
	//			}
	//		}else {
	//			log.Info("Chưa có ai booking " + booking.TimeBooking)
	//			successTimeBooking = booking.TimeBooking
	//		}
	//	}
	//}else {
	//	log.Info("Chưa có ai booking " + booking.TimeBooking)
	//	successTimeBooking = booking.TimeBooking
	//}

	//Không check  thời gian này đã được book hay chưa
	successTimeBooking := booking.TimeBooking
	////////////////////

	conctRemote.SetImplicitWaitTimeout(time.Second * 60)
	elemTimeBooking, err := conctRemote.FindElement(selenium.ByXPATH,
		"//span[contains(text(),'"+successTimeBooking+"')]")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemTimeBooking.Click()
	log.Printf(booking.Username + " Đã đã chọn giờ " + successTimeBooking)

	elemCheckboxAgree, err := conctRemote.FindElement(selenium.ByCSSSelector, "label[for='cbxAgreePolicy']")
	if err != nil {
		log.Error(err.Error())
		return err
	}else {
		elemCheckboxAgree.Click()
	}
	log.Printf(booking.Username + " Đã chọn đồng ý các điều khoản")

	var ButtonSave string
	elemSaveBooking, err := conctRemote.FindElement(selenium.ByCSSSelector,
		"button.btn.btn-primary.mr-3[type='submit']")
	if err != nil {
		log.Error(err.Error())
		return err
	} else {
		ButtonSave, _ = elemSaveBooking.Text()
		log.Printf(booking.Username + " Đã nhận diện được nút "+ ButtonSave)
		elemSaveBooking.Click()

		//update Db
		var updateNextDayBooking int
		hourCurrent := time.Now().Hour()
		if hourCurrent == 23 {
			updateNextDayBooking = time.Now().AddDate(0, 0, +15).Day()

		} else {
			updateNextDayBooking = time.Now().AddDate(0, 0, +14).Day()
		}

		updateBooking := model.Booking{
			UserId:           booking.UserId,
			DateBooking: strconv.Itoa(updateNextDayBooking),
			CreatedBooking: booking.CreatedBooking,
			CompletedBooking:   time.Now(),
		}
		log.Printf("Tự động cập nhật ngày "+strconv.Itoa(updateNextDayBooking)+ " cho " + booking.Username)
		_, err = userBooking.UpdateUsersBooking(ctx, updateBooking)
		if err != nil {
			log.Error(err.Error())
			return err
		}
	}

	shot, err := conctRemote.Screenshot()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	nameImage := booking.Username+"-Booking-"+booking.Amenity+"-"+time.Now().Format("20060102150405")
	err = helper.HelpSaveImage(shot, nameImage)
	if err != nil {
		log.Error(err.Error())
		return err
	}else {
		locateImage := "./log_files/screenshots/" + nameImage + ".png"
		err = sendMail.HelperSendMailBookingSuccess(booking, locateImage, successTimeBooking)
		if err != nil {
			log.Error(err.Error())
			return err
		}
	}
	log.Printf(booking.Username + " Đã Hoàn thành Booking và gửi Mail "+ ButtonSave)

	return nil
}


// Booking Interval

func (s *Selenium) BookingTennisInterval(booking model.Booking, userBooking repository.UserBooking, sendMail helper.SendMail) error {
	ctx := context.Background()
	conctRemote, err := s.getRemote()
	if err != nil {
		log.Error(err.Error()+">>>"+booking.Username)
		return err
	}
	defer conctRemote.Quit()

	err = conctRemote.Get("https://vn.propertycube.asia/account")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	conctRemote.SetImplicitWaitTimeout(time.Second * 60)

	if title, err := conctRemote.Title(); err == nil {
		fmt.Printf("Page title: %s\n", title)
	} else {
		log.Error(err.Error())
		return err
	}

	conctRemote.ResizeWindow("note", 1080, 720)

	elemUserNameOrEmailAddress, err := conctRemote.FindElement(selenium.ByName, "userNameOrEmailAddress")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemUserNameOrEmailAddress.SendKeys(booking.Username)
	log.Printf(booking.Username + "Đã điền được User")

	elemPassword, err := conctRemote.FindElement(selenium.ByName, "password")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemPassword.SendKeys(booking.Password)
	log.Printf(booking.Username + " Đã điền được Password")

	elemBtLogin, err := conctRemote.FindElement(selenium.ByCSSSelector, "button[type=\"submit\"]")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemBtLogin.Click()
	log.Printf(booking.Username + " Đã đã nhấn login")

	elemTenancy, err := conctRemote.FindElement(selenium.ByCSSSelector, "[title='MASTERI']")
	if err != nil {
		log.Info(err.Error())
		log.Info(booking.Username + "Chỉ có Thuê 1 căn")
	}else {
		elemTenancy.Click()
		log.Printf(booking.Username + " Đã nhấn vào Tenancy MASTERI")
	}


	elemUserProfile, err := conctRemote.FindElement(selenium.ByCSSSelector, "img.rounded.ng-star-inserted")
	if err != nil {
		log.Error(err.Error())
		log.Error(booking.Username + " login thất bại")
		return err
	}
	elemUserProfile.Click()
	log.Printf(booking.Username + " Đã nhấn vào Picture Profile")

	elemCheckUserProfile, err := conctRemote.FindElement(selenium.ByCSSSelector, ".media-body.text-light")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	name, _ := elemCheckUserProfile.Text()
	log.Printf(booking.Username + " so sánh với "+ name)

	elemCheckLanguage, err := conctRemote.FindElement(selenium.ByCSSSelector, ".ml-2.hide-mobile")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	language, _ := elemCheckLanguage.Text()
	if language != "English" {
		elemChangeLanguage, err := conctRemote.FindElement(selenium.ByCSSSelector, ".flag-icon.flag-icon-vn")
		if err != nil {
			log.Error(err.Error())
			return err
		}
		elemChangeLanguage.Click()
		log.Printf(booking.Username + " chọn Change language")

		elemChooseEnglish, err := conctRemote.FindElement(selenium.ByCSSSelector, ".flag-icon.flag-icon-gb")
		if err != nil {
			log.Error(err.Error())
			return err
		}
		elemChooseEnglish.Click()
		log.Printf(booking.Username + " chọn language English")
	}

	if booking.Homes == "yes" {
		elemHomes, err :=  conctRemote.FindElement(selenium.ByCSSSelector, "span.ng-star-inserted")
		if err != nil {
			log.Error(err.Error())
			return err
		}
		elemHomes.Click()
		log.Printf(booking.Username + " chọn Homes")

		HomeId, err :=  conctRemote.FindElement(selenium.ByXPATH, "//li[contains(text(),'"+ booking.HomeID+"')]")
		if err != nil {
			log.Error(err.Error())
			return err
		}
		HomeId.Click()
		log.Printf(booking.Username + " chọn Homes")
	}


	time.Sleep(time.Second * 5)
	conctRemote.SetImplicitWaitTimeout(time.Second * 60)

	elemBooking, err := conctRemote.FindElement(selenium.ByCSSSelector, "#menuItemFacilityManagement " +
		"> li.dropdown.ng-star-inserted > a > span.menu-name")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemBooking.Click()
	log.Printf(booking.Username + " Đã nhấn vào Booking cha")

	elemBookingChild, err := conctRemote.FindElement(selenium.ByCSSSelector, "#menuItemFacilityManagement " +
		"> li.dropdown.ng-star-inserted > ul > li:nth-child(1) > a > span.menu-name")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemBookingChild.Click()
	log.Printf(booking.Username + " Đã nhấn vào Booking con")

	var amenities string
	var amenity string
	switch booking.Amenities {
	case "TENNIS NEW":
		amenities = "a[href='#amenityGroup14']"
		if booking.Amenity == "Tennis 1 New" {
			amenity = "label[for=\"amenities62\"]"
		}else{
			amenity = "label[for=\"amenities64\"]"
		}
	}

	elemAmenities, err := conctRemote.FindElement(selenium.ByCSSSelector, amenities)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemAmenities.Click()
	log.Printf(booking.Username + " Đã nhấn vào Amenities " + booking.Amenities)

	elemAmenity, err := conctRemote.FindElement(selenium.ByCSSSelector, amenity)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemAmenity.Click()
	log.Printf(booking.Username + " Đã đã nhấn vào Amenity " + booking.Amenity)

	conctRemote.SetImplicitWaitTimeout(time.Second * 10)
	elemDateBooking, err := conctRemote.FindElement(selenium.ByXPATH,
		"//span[text()='"+booking.DateBooking+"']")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemDateBooking.Click()
	log.Printf(booking.Username + " Đã chọn ngày " + booking.DateBooking)

	elemGetTimesUnavailable, err := conctRemote.FindElements(selenium.ByCSSSelector, "td.fc-event-container a.unavailable")
	if err != nil {
		log.Error(err.Error())
		return err
	}

	var successTimeBooking string
	if len(elemGetTimesUnavailable) > 0 {
		for _, timeUnavailable := range elemGetTimesUnavailable {
			timeU, err := timeUnavailable.Text()
			if err != nil {
				log.Error(err.Error())
				return err
			}

			if timeU == booking.TimeBooking {
				log.Error(booking.Username + " không thể chọn " + booking.TimeBooking)
				//check xem là đã có session nào book thành công chưa?
				checkUserBooked := model.Booking{
					Username:          booking.Username,
					DateBooking:       booking.DateBooking,
					Amenity:           booking.Amenity,
				}

				_, err := userBooking.CheckUsersBooked(ctx, checkUserBooked)
				if err != nil {
					log.Printf(booking.Username + " đã booking thành công " + booking.TimeBooking)
					log.Error(err.Error())
					return err
				}else {
					log.Printf(booking.Username + " Đã chưa booking thành công " + booking.TimeBooking +
						" Chuyển sang " + booking.TimeBookingBackup)
					successTimeBooking = booking.TimeBookingBackup
					break
				}
			}else {
				log.Info("Chưa có ai booking " + booking.TimeBooking)
				successTimeBooking = booking.TimeBooking
			}
		}
	}else {
		log.Info("Chưa có ai booking " + booking.TimeBooking)
		successTimeBooking = booking.TimeBooking
	}

	conctRemote.SetImplicitWaitTimeout(time.Second * 60)
	elemTimeBooking, err := conctRemote.FindElement(selenium.ByXPATH,
		"//span[contains(text(),'"+successTimeBooking+"')]")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	elemTimeBooking.Click()
	log.Printf(booking.Username + " Đã đã chọn giờ " + successTimeBooking)

	elemCheckboxAgree, err := conctRemote.FindElement(selenium.ByCSSSelector, "label[for='cbxAgreePolicy']")
	if err != nil {
		log.Error(err.Error())
		return err
	}else {
		elemCheckboxAgree.Click()

		var updateNextDayBooking int
		hourCurrent := time.Now().Hour()
		if hourCurrent == 23 {
			updateNextDayBooking = time.Now().AddDate(0, 0, +15).Day()

		} else {
			updateNextDayBooking = time.Now().AddDate(0, 0, +14).Day()
		}

		updateBooking := model.Booking{
			UserId:           booking.UserId,
			DateBooking: strconv.Itoa(updateNextDayBooking),
			CreatedBooking: booking.CreatedBooking,
			CompletedBooking:   time.Now(),
		}
		log.Printf("Tự động cập nhật ngày "+strconv.Itoa(updateNextDayBooking)+ " cho " + booking.Username)
		_, err = userBooking.UpdateUsersBooking(ctx, updateBooking)
		if err != nil {
			log.Error(err.Error())
			return err
		}
	}

	log.Printf(booking.Username + " Đã chọn đồng ý các điều khoản")

	shot, err := conctRemote.Screenshot()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	elemSaveBooking, err := conctRemote.FindElement(selenium.ByCSSSelector,
		"button.btn.btn-primary.mr-3[type='submit']")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	ButtonSave, _ := elemSaveBooking.Text()
	log.Printf(booking.Username + " Đã nhận diện được nút "+ ButtonSave)
	elemSaveBooking.Click()

	nameImage := booking.Username+"-Booking-"+booking.Amenity+"-"+time.Now().Format("20060102150405")
	err = helper.HelpSaveImage(shot, nameImage)
	if err != nil {
		log.Error(err.Error())
		return err
	}else {
		locateImage := "./log_files/screenshots/" + nameImage + ".png"
		err = sendMail.HelperSendMailBookingSuccess(booking, locateImage, successTimeBooking)
		if err != nil {
			log.Error(err.Error())
			return err
		}
	}
	log.Printf(booking.Username + " Đã Hoàn thành Booking và gửi Mail "+ ButtonSave)
	return nil
}