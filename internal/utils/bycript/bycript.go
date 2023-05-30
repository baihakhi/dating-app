package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Encrypt(pass string) (string, error) {
	hashedPassword, err := Hash(pass)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
