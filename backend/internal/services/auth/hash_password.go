package services_auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (service *AuthService) HashPassword(password string) (string, error) {
	password_bytes := []byte(password)
	// Hashing the password with the default cost of 10
	hashed_password, err := bcrypt.GenerateFromPassword(password_bytes, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	fmt.Println(string(hashed_password))
	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(hashed_password, password)
	fmt.Println(err) // nil means it is a match

	return string(hashed_password), nil
}
