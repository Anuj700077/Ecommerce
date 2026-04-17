package cart

import (
	"Ecommerce/cart"
	"errors"
)

func GetCartService(userID uint) ([]cart.Cart, int, error) {

	cartItems, err := cart.GetCartByUserID(userID)
	if err != nil {
		return nil, 0, err
	}

	if len(cartItems) == 0 {
		return nil, 0, errors.New("cart is empty")
	}


	grandTotal := 0
	for _, item := range cartItems {
		grandTotal += item.TotalPrice
	}

	return cartItems, grandTotal, nil
}
