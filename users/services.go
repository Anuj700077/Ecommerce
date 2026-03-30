package Users

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/mail"
	"time"
)

var otpStore = make(map[string]string)

var (
	ErrEmailRequired = errors.New("EMAIL_REQUIRED")
	ErrInvalidEmail  = errors.New("INVALID_EMAIL_FORMAT")
	ErrInvalidOTP    = errors.New("INVALID_OTP")
	ErrFieldsMissing = errors.New("FIELDS_MISSING")
)

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func SendOTPService(email string) error {

	if email == "" {
		return ErrEmailRequired
	}

	if !isValidEmail(email) {
		return ErrInvalidEmail
	}

	otp := generateOTP()
	otpStore[email] = otp

	log.Println("OTP for", email, "is:", otp)

	return nil
}

func verifyOTP(email, otp string) bool {
	storedOTP, exists := otpStore[email]
	if !exists {
		return false
	}
	return storedOTP == otp
}

func VerifyOTPAndRegister(user *User, otp string) error {

	if user.Name == "" ||
		user.Email == "" ||
		user.Phone == "" ||
		user.Password == "" {
		return ErrFieldsMissing
	}

	if !isValidEmail(user.Email) {
		return ErrInvalidEmail
	}

	if !verifyOTP(user.Email, otp) {
		return ErrInvalidOTP
	}

	// Save user
	err := CreateUser(user)
	if err != nil {
		return err
	}

	// Here...Remove OTP after success
	delete(otpStore, user.Email)

	return nil
}
