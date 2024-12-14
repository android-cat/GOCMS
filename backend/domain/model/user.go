package model

import (
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)
type User struct {
	gorm.Model
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Name	string	`json:"name"`
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}
