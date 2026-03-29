package Users

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Request for sending OTP
type OTPRequest struct {
	Email string `json:"email"`
}

// Request for verifying OTP + register
type VerifyRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	OTP      string `json:"otp"`
}
