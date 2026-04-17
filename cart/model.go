package cart

import "time"

type Cart struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	ProductID    uint      `json:"product_id"`
	ProductName  string    `json:"product_name"`
	ProductPrice int       `json:"product_price"`
	Quantity     int       `json:"quantity"`
	TotalPrice   int       `json:"total_price"`
	CreatedAt    time.Time `json:"created_at"`
}
