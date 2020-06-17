package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PasswordEncoder struct{}

func NewPasswordEncoder() *PasswordEncoder {
	return &PasswordEncoder{}
}

func (enc PasswordEncoder) encodePassword(password string) (string, error) {
	bytePwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("Cant encode password: %w", err)
	}
	return string(hash), nil
}

func (enc PasswordEncoder) isPasswordValid(hashPwd, pwd string) bool {
	byteHashPwd := []byte(hashPwd)
	bytePwd := []byte(pwd)
	err := bcrypt.CompareHashAndPassword(byteHashPwd, bytePwd)
	return err == nil
}
