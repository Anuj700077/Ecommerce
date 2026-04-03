package Products

import (
	"Ecommerce/Products"
	"errors"
)

var (
	ErrInvalidID       = errors.New("INVALID_ID")
	ErrProductNotFound = errors.New("PRODUCT_NOT_FOUND")
)

func DeleteProducts(id int64) error {
	if id <= 0 {
		return ErrInvalidID
	}
	err := Products.DeleteProduct(id)

	if err != nil {
		if err.Error() == "PRODCUT_NOT_FOUND" {
			return ErrProductNotFound
		}
		return err
	}
	return nil
}
