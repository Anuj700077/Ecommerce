package Products

import (
	"Ecommerce/database"
	"database/sql"
	"errors"
)

func GetAllProducts() ([]Product, error) {

	
	if database.SQLDB == nil {
		return nil, errors.New("database not connected")
	}

	query := `SELECT id, name, description, price FROM products`

	rows, err := database.SQLDB.Query(query) 
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if len(products) == 0 {
		return nil, sql.ErrNoRows
	}
	return products, nil
}
