package cart

import (
	"Ecommerce/cart"
	"Ecommerce/database"
	"errors"
)

func getProductDetails(productID uint) (string, int, error) {

	query := `SELECT name, price FROM products WHERE id=$1`

	var name string
	var price int

	err := database.SQLDB.QueryRow(query, productID).Scan(&name, &price)
	if err != nil {
		return "", 0, errors.New("product not found")
	}

	return name, price, nil
}

func AddToCartService(userID uint, productID uint, quantity int) error {

	if productID == 0 {
		return errors.New("product_id is required")
	}

	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	productName, productPrice, err := getProductDetails(productID)
	if err != nil {
		return err
	}

	totalPrice := quantity * productPrice

	cartItem := &cart.Cart{
		UserID:       userID,
		ProductID:    productID,
		ProductName:  productName,
		ProductPrice: productPrice,
		Quantity:     quantity,
		TotalPrice:   totalPrice,
	}

	return cart.CreateCartItem(cartItem)
}


