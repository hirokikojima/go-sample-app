package authorizer

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (*string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	encrypted := string(bytes)
	return &encrypted, nil
}

func ComparePassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}