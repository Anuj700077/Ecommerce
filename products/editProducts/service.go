package Products

import (
	"Ecommerce/Products"
	"errors"
)

var (
	ErrInvalidProduct  = errors.New("INVALID_PRODUCT_DATA")
	ErrProductNotFound = errors.New("PRODUCT_NOT_FOUND")
)

func UpdateProductService(product *Products.Product) error {

	if product.ID == 0 || product.Name == "" || product.Price <= 0 {
		return ErrInvalidProduct
	}

	err := Products.UpdateProduct(product)

	if err != nil {
		if err.Error() == "PRODUCT_NOT_FOUND" {
			return ErrProductNotFound
		}
		return err
	}

	return nil
}
