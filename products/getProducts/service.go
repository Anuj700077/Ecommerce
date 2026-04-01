package Products

import (
	"Ecommerce/Products"
	"errors"
)


var ErrNoProductsFound = errors.New("NO_PRODUCTS")

func GetProduct() ([]Products.Product, error) {

	products, err := Products.GetProducts()
	if err != nil {
		return nil, err
	}

	
	if len(products) == 0 {
		return nil, ErrNoProductsFound
	}

	return products, nil
}
