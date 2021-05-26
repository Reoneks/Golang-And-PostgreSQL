package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

func Encrypt(password string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(password))
	if err != nil {
		return "", err
	}
	sha := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return sha, nil
}

func Compare(userpass, password string) error {
	userPassword, err := Encrypt(password)
	if err != nil {
		return err
	}
	if userpass != userPassword {
		return errors.New("entered the wrong password")
	}
	return nil
}
