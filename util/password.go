package util

import "golang.org/x/crypto/bcrypt"

func HashedPassword(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedByte), nil
}
