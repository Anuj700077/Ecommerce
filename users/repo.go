package users

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

	err := database.DB.Create(user).Error
	if err != nil {
		log.Println("Error inserting user:", err)
		return err
	}

	return nil
}
