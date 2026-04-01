package otpandregister

import (
	"Ecommerce/Users"
	"Ecommerce/utils"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/mail"
	"time"
)

var otpStore = make(map[string]Users.OTPData)
var (
	ErrEmailRequired   = errors.New("EMAIL_REQUIRED")
	ErrInvalidEmail    = errors.New("INVALID_EMAIL_FORMAT")
	ErrOTPExpired      = errors.New("OTP_EXPIRED")
	ErrOTPNotFound     = errors.New("OTP_NOT_FOUND")
	ErrInvalidOTP      = errors.New("INVALID_OTP")
	ErrFieldsMissing   = errors.New("FIELDS_MISSING")
	ErrUserNotFound    = errors.New("USER_NOT_FOUND")
	ErrInvalidPassword = errors.New("INVALID_PASSWORD")
)

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// This function is for Generate OTP
func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// Here in this function The OTP will send
func SendOTPService(email string) error {

	if email == "" {
		return ErrEmailRequired
	}

	if !isValidEmail(email) {
		return ErrInvalidEmail
	}

	otp := generateOTP()

	otpStore[email] = Users.OTPData{
		Code:      otp,
		ExpiresAt: time.Now().Add(40 * time.Second), //  expiry time of OTP
	}

	log.Println("OTP for", email, "is:", otp)

	return nil
}

// This function is for Verify the  OTP
func verifyOTP(email, otp string) error {

	data, exists := otpStore[email]
	if !exists {
		return ErrOTPNotFound
	}

	//here it checks the Expiry
	if time.Now().After(data.ExpiresAt) {
		delete(otpStore, email)
		return ErrOTPExpired
	}

	// it is for Wrong OTP
	if data.Code != otp {
		return ErrInvalidOTP
	}

	return nil
}
func VerifyOTPAndRegister(user *Users.User, otp string) error {

	if user.Name == "" ||
		user.Email == "" ||
		user.Phone == "" ||
		user.Password == "" {
		return ErrFieldsMissing
	}

	if !isValidEmail(user.Email) {
		return ErrInvalidEmail
	}

	// here it Verify the  OTP
	err := verifyOTP(user.Email, otp)
	if err != nil {
		return err
	}

	//In this line password hashing technique done
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	err = Users.CreateUser(user)

	// Remove OTP after success
	delete(otpStore, user.Email)

	return nil
}
