package Users

import (
	"Ecommerce/utils"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/mail"
	"time"
)

// In this variables In-memory OTP store
var otpStore = make(map[string]OTPData)

// These are the Custom Errors
var (
	ErrEmailRequired   = errors.New("EMAIL_REQUIRED")
	ErrInvalidEmail    = errors.New("INVALID_EMAIL_FORMAT")
	ErrInvalidOTP      = errors.New("INVALID_OTP")
	ErrFieldsMissing   = errors.New("FIELDS_MISSING")
	ErrOTPExpired      = errors.New("OTP_EXPIRED")
	ErrOTPNotFound     = errors.New("OTP_NOT_FOUND")
	ErrUserNotFound    = errors.New("USER_NOT_FOUND")
	ErrInvalidPassword = errors.New("INVALID_PASSWORD")
)

// this function is for Email validation
func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// This function is for Generate OTP
func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// Here in this The OTP will send
func SendOTPService(email string) error {

	if email == "" {
		return ErrEmailRequired
	}

	if !isValidEmail(email) {
		return ErrInvalidEmail
	}

	otp := generateOTP()

	otpStore[email] = OTPData{
		Code:      otp,
		ExpiresAt: time.Now().Add(30 * time.Second), // ⏳ expiry
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

// This fucntion is created for verify and register the user
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

	// here it Verify the  OTP
	err := verifyOTP(user.Email, otp)
	if err != nil {
		return err
	}

	
	// Hash password before saving
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	err = CreateUser(user)

	// Remove OTP after success
	delete(otpStore, user.Email)

	return nil
}

func LoginService(email, password string) (*User, error) {

	if email == "" || password == "" {
		return nil, ErrFieldsMissing
	}

	if !isValidEmail(email) {
		return nil, ErrInvalidEmail
	}

	// check user exists
	user, err := GetUserByEmail(email)
	if err != nil {
		return nil, ErrUserNotFound
	}

	// check password (simple compare for now)
	if !utils.CheckPasswordHash(password, user.Password) {
	return nil, ErrInvalidPassword
}


	return user, nil
}
