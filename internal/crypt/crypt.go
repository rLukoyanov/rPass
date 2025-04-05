package crypt

import (
	"crypto/sha256"
	"fmt"
	"syscall"

	"golang.org/x/term"
)

func GetMasterKey() ([]byte, error) {
	fmt.Print("Введите мастер пароль")
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(password)
	return hash[:], nil
}
