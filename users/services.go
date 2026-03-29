package Users

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"time"
)

// In-memory OTP store
var otpStore = make(map[string]string)

// Generate OTP
func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// Send Email using SMTP
func sendEmail(to, otp string) error {

	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	if from == "" || password == "" {
		return errors.New("email credentials not set in .env")
	}

	message := []byte(
		"Subject: OTP Verification\n\n" +
			"Your OTP is: " + otp,
	)

	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		from,
		[]string{to},
		message,
	)

	if err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	return nil
}

// Send OTP
func SendOTPService(email string) error {

	if email == "" {
		return errors.New("email is required")
	}

	otp := generateOTP()

	// Save OTP
	otpStore[email] = otp

	// Send email
	log.Println("OTP for", email, "is:", otp)

	// Also print OTP (for testing)
	log.Println("OTP:", otp)

	return nil
}

// Verify OTP
func verifyOTP(email, otp string) bool {
	storedOTP, exists := otpStore[email]
	if !exists {
		return false
	}
	return storedOTP == otp
}

// Verify OTP + Register
func VerifyOTPAndRegister(user *User, otp string) error {

	if user.Name == "" ||
		user.Email == "" ||
		user.Phone == "" ||
		user.Password == "" {
		return errors.New("all fields are required")
	}

	if !verifyOTP(user.Email, otp) {
		return errors.New("invalid OTP")
	}

	// Save user
	err := CreateUser(user)
	if err != nil {
		return err
	}

	// Remove OTP after success
	delete(otpStore, user.Email)

	return nil
}
