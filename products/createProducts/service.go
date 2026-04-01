package Products

import (
	"Ecommerce/Products"
	"errors"
)

var (
	ErrInvalidProduct = errors.New("INVALID_PRODUCT_DATA")
)

func CreateProductService(product *Products.Product) error {

	if product.Name == "" || product.Price <= 0 {
		return ErrInvalidProduct
	}

	return Products.CreateProduct(product)
}
