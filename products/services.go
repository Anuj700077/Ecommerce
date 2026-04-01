package Products

import (
	"errors"
)

var (
	ErrInvalidProduct = errors.New("INVALID_PRODUCT_DATA")
)


func CreateProductService(product *Product) error {

	if product.Name == "" || product.Price <= 0 {
		return ErrInvalidProduct
	}

	return CreateProduct(product)
}

var ErrNoProductsFound = errors.New("NO_PRODUCTS")

func GetProduct() ([]Product, error) {

	products, err := GetProducts()
	if err != nil {
		return nil, err
	}

	
	if len(products) == 0 {
		return nil, ErrNoProductsFound
	}

	return products, nil
}
