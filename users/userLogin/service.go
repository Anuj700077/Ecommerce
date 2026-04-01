package userlogin

import (
	"Ecommerce/Users"
	"Ecommerce/utils"
	"errors"
	"net/mail"
)

var (
	ErrFieldsMissing   = errors.New("FIELDS_MISSING")
	ErrInvalidEmail    = errors.New("INVALID_EMAIL_FORMAT")
	ErrUserNotFound    = errors.New("USER_NOT_FOUND")
	ErrInvalidPassword = errors.New("INVALID_PASSWORD")
)

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func LoginService(email, password string) (*Users.User, error) {

	if email == "" || password == "" {
		return nil, ErrFieldsMissing
	}

	if !isValidEmail(email) {
		return nil, ErrInvalidEmail
	}

	// check user exists
	user, err := Users.GetUserByEmail(email)

	if err != nil {
		return nil, ErrUserNotFound
	}

	// check password (simple compare for now)
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, ErrInvalidPassword
	}

	return user, nil
}
