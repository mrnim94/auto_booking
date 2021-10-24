package helper

import (
	"auto-booking/log"
	"auto-booking/model"
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type AccountMail struct {
	Host string
	Port int
	Username string
	Password string
}

type SendMail interface {
	HelperSendMailBookingSuccess(booking model.Booking, locateScreenShot string, successTimeBooking string) error
}

func NewSendMail(ac *AccountMail) SendMail {
	return &AccountMail{
		Host:     ac.Host,
		Port:     ac.Port,
		Username: ac.Username,
		Password: ac.Password,
	}
}

func (ac *AccountMail) connectMailServer() *gomail.Dialer {
	d := gomail.NewDialer(ac.Host, ac.Port, ac.Username, ac.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return d
}

func (ac *AccountMail) HelperSendMailBookingSuccess(booking model.Booking, locateScreenShot string, successTimeBooking string) error {
	con := ac.connectMailServer()
	//monthBooking := time.Now().Month()
	bodyMessageNim :=
		"<p>Hello: <strong>"+booking.Username+"</strong></p>" +
		"\n<p>He thong da Booking Tennis vao ngay <strong>"+booking.DateBooking+"</strong>, khung gio <strong>"+successTimeBooking+"</strong></p>" +
		"\n<p>Moi thong tin chi tiet trong file dinh kem.</p>" +
		"\n<p>Merci beaucoup!</p>"

	m := gomail.NewMessage()
	m.SetHeader("From", ac.Username)
	m.SetHeader("To", "mr.nim94@gmail.com")
	//m.SetHeader("To", booking.Username)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", bodyMessageNim)
	m.Attach(locateScreenShot)

	// Send the email to Bob, Cora and Dan.
	if err := con.DialAndSend(m); err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}