package Users

import "time"

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Here...Request for sending OTP
type OTPRequest struct {
	Email string `json:"email"`
}

// Here...Request for verifying OTP + register
type VerifyRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	OTP      string `json:"otp"`
}

//Here OTp expire
type OTPData struct {
	Code      string
	ExpiresAt time.Time
}

// Login Request
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
