package address

import "time"

type Address struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	Address1  string    `json:"address1"`
	Pincode   string    `json:"pincode"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
}
