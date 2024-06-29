package helper

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type HashingHelper interface {
	HashPassword(pwd string) (string, error)
	CheckPassword(pwd string, hash []byte) (bool, error)
}

type HashImplementation struct {
}

func (h *HashImplementation) HashPassword(pwd string) (string, error) {
	costInt, err := strconv.Atoi(os.Getenv("COST"))
	if err != nil {
		return "", err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), costInt)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (h *HashImplementation) CheckPassword(pwd string, hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, []byte(pwd))
	if err != nil {
		return false, err
	}
	return true, nil
}