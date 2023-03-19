package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordManager interface {
	HashPassword(password string) (string, error)
	IsValidPassword(password, hash string) bool
}

func NewPasswordManager() PasswordManager {
	return &passwordManager{}
}

type passwordManager struct{}

func (p *passwordManager) HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func (p *passwordManager) IsValidPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
