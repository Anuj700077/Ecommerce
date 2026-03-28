package users

import "errors"

func RegisterUser(user *User) error {

	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Phone == "" {
		return errors.New("phone is required")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	return CreateUser(user)
}
