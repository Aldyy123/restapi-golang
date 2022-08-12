package controllers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	return string(hash), err
}

func ValidationPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil

}
