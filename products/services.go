package Products

import (
	"database/sql"
	"errors"
)





var ErrNoProductsFound = errors.New("NO_PRODUCTS")

func GetProductsService() ([]Product, error) {

	products, err := GetAllProducts()
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, ErrNoProductsFound
		}
		return nil, err
	}
	return products, nil
}
