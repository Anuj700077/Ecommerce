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

	//  IMPORTANT: set default role if empty
	if user.Role == "" {
		user.Role = "user"
	}

	query := `
	INSERT INTO users (name, email, phone, password, role)
	VALUES ($1, $2, $3, $4, $5)
	`

	err := database.DB.Exec(query,
		user.Name,
		user.Email,
		user.Phone,
		user.Password,
		user.Role,
	).Error

	if err != nil {
		log.Println("Error inserting user:", err)
		return err
	}

	return nil
}

func GetUserByEmail(email string) (*User, error) {

	if database.SQLDB == nil {
		return nil, errors.New("database not connected")
	}

	query := `SELECT id, name, email, phone, password, role FROM users WHERE email=$1`

	row := database.SQLDB.QueryRow(query, email)

	var user User

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.Password,
		&user.Role,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
