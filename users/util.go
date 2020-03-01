package users

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/argon2"
	"os"
)

func convertFileToUser(file *os.File) (*UserInfo, error) {
	userInfo := &UserInfo{}
	decoder := json.NewDecoder(file)
	err := decoder.Decode(userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func convertUserToFileContent(user *UserInfo) (string, error) {
	content, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func generatePassword(password string, salt []byte) ([]byte) {
	memory := uint32(64 * 1024)
	iterations := uint32(3)
	parallelism := uint8(2)
	keyLength := uint32(32)

	hash := argon2.IDKey(
		[]byte(password), salt, iterations, memory, parallelism, keyLength)
	return hash
}

func checkPassword(entered string, password []byte, salt []byte) error {
	checked := generatePassword(entered, salt)
	if bytes.Compare(checked, password) != 0 {
		return errors.New("password does not match")
	}
	return nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

