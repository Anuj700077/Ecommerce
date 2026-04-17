package cart

import "Ecommerce/database"

func CreateCartItem(c *Cart) error {

	query := `
	INSERT INTO cart (user_id, product_id, product_name, product_price, quantity, total_price)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	return database.DB.Exec(query,
		c.UserID,
		c.ProductID,
		c.ProductName,
		c.ProductPrice,
		c.Quantity,
		c.TotalPrice,
	).Error
}

// 🔹 Get all cart items for a user
func GetCartByUserID(userID uint) ([]Cart, error) {

	query := `
	SELECT id, user_id, product_id, product_name, product_price, quantity, total_price, created_at
	FROM cart
	WHERE user_id=$1
	ORDER BY created_at DESC
	`

	rows, err := database.SQLDB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartItems []Cart

	for rows.Next() {
		var c Cart

		err := rows.Scan(
			&c.ID,
			&c.UserID,
			&c.ProductID,
			&c.ProductName,
			&c.ProductPrice,
			&c.Quantity,
			&c.TotalPrice,
			&c.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		cartItems = append(cartItems, c)
	}

	return cartItems, nil
}
