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
		product.Name, product.Description, product.Price,
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
			&p.ID, &p.Name, &p.Description, &p.Price, p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func UpdateProduct(product *Product) error {

	if database.SQLDB == nil {
		return errors.New("database not connected")
	}
	query := `
	UPDATE products
	SET name = $1, description = $2, price = $3
	WHERE id = $4
	`

	result, err := database.SQLDB.Exec(query,
		product.Name, product.Description, product.Price, product.ID,
	)

	if err != nil {
		log.Println("Error updating product:", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("PRODUCT_NOT_FOUND")
	}
	return nil
}

func DeleteProduct(id int64) error {

	if database.SQLDB == nil {
		return errors.New("database not connected")
	}

	query := `
	DELETE FROM products
	WHERE id = $1
	`

	result, err := database.SQLDB.Exec(query, id)
	if err != nil {
		log.Println("Error deleting product:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("PRODUCT_NOT_FOUND")
	}

	return nil
}
