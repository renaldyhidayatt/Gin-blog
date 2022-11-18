package utils

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	pw := []byte(password)

	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return string(result)
}

func ComparePassword(hashpassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashpassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)

	return err
}
