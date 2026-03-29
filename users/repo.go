package Users

import (
	"errors"
	"log"

	"Ecommerce/database"
)

func CreateUser(user *User) error {

	if database.DB == nil {
		return errors.New("database not connected")
	}

	if user == nil {
		return errors.New("user data is empty")
	}

	// Raw SQL Insert Query
	query := `
	INSERT INTO users (name, email, phone, password)
	VALUES (?, ?, ?, ?)
	`

	err := database.DB.Exec(query,
		user.Name,
		user.Email,
		user.Phone,
		user.Password,
	).Error

	if err != nil {
		log.Println("Error inserting user:", err)
		return err
	}

	return nil
}
