package main

import "fmt"

// ---

type OTPService interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

// ---

type OTP struct {
	service OTPService
}

func (o *OTP) genAndSendOTP(otpLength int) error {
	otp := o.service.genRandomOTP(otpLength)
	o.service.saveOTPCache(otp)
	message := o.service.getMessage(otp)
	return o.service.sendNotification(message)
}

// ---

type SMS struct {
	OTP
}

func (s *SMS) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *SMS) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *SMS) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *SMS) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

// ---

type Email struct {
	OTP
}

func (s *Email) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (s *Email) getMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (s *Email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}

// ---

func main() {
	smsOTP := &SMS{}
	o := OTP{service: smsOTP}
	o.genAndSendOTP(4)

	fmt.Println()

	emailOTP := &Email{}
	o = OTP{service: emailOTP}
	o.genAndSendOTP(4)
}
