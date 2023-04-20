package encryption

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(passwordHashed), nil
}

func ValidatePassword(hashed, pass string) error {
	fmt.Println("hashed: ", hashed, "pass: ", pass)
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass)); err != nil {
		return err
	}
	return nil
}
