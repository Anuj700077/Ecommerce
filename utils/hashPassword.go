package utils

import "golang.org/x/crypto/bcrypt"

// Hash Password
func HashPassword(password string) (string, error) {

	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// Compare Password
func CheckPasswordHash(password, hashedPassword string) bool {

	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)

	return err == nil
}
