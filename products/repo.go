package Products

import (
	"Ecommerce/database"
	"errors"
	"log"
)

func CreateProduct(product *Product) error {

	if database.SQLDB == nil {
		return errors.New("database not connected")
	}

	query := `
	INSERT INTO products (name, description, price)
	VALUES ($1, $2, $3)
	`

	_, err := database.SQLDB.Exec(query,
		product.Name,
		product.Description,
		product.Price,
	)

	if err != nil {
		log.Println("Error inserting product:", err)
		return err
	}

	return nil
}

func GetProducts() ([]Product, error) {

	if database.SQLDB == nil {
		return nil, errors.New("database not connected")
	}

	query := `
	SELECT id, name, description, price, created_at
	FROM products
	ORDER BY id DESC
	`

	rows, err := database.SQLDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
