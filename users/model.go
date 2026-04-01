package Users

import "time"

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

// OTP Request
type OTPRequest struct {
	Email string `json:"email"`
}

// Verify + Register
type VerifyRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	OTP      string `json:"otp"`
}

// OTP store
type OTPData struct {
	Code      string
	ExpiresAt time.Time
}

// Login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
