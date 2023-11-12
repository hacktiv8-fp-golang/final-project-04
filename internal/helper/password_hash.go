package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) (string, Error) {
	salt := 12
	password := []byte(p)
	hashedPass, err := bcrypt.GenerateFromPassword(password, salt)

	if err != nil {
		return "", InternalServerError("Failed to hash the password")
	}

	return string(hashedPass), nil
}

func ComparePassword(p string, h string) bool {
	password, hashedPassword := []byte(p), []byte(h)

	err := bcrypt.CompareHashAndPassword(hashedPassword, password)

	return err == nil
}
