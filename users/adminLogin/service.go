package adminlogin

import (
	"Ecommerce/utils"
	"errors"
)
var (
	AdminEmail    = "admin@gmail.com"
	AdminPassword = "$2a$10$7eX.tP7T23lUgFQc6FGute9BC4DAuKtoAa4x9HUaejk6rZuTDfooK"
	// password = "admin123"
)

var (
	ErrFieldsMissing = errors.New("FIELDS_MISSING")
	ErrAdminInvalidCredentials = errors.New("ADMIN_INVALID_CREDENTIALS")
)

// Admin login credentials
func AdminLoginService(email, password string) error {

	if email == "" || password == "" {
		return ErrFieldsMissing
	}

	if email != AdminEmail {
		return ErrAdminInvalidCredentials
	}

	// compare hashed password
	if !utils.CheckPasswordHash(password, AdminPassword) {
		return ErrAdminInvalidCredentials
	}

	return nil
}
