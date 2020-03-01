package util

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/marcsj/jotfly-server/users"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/argon2"
	"os"
	"time"
)

const saltLength = 16

func convertFileToUser(file *os.File) (*users.UserInfo, error) {
	userInfo := &users.UserInfo{}
	decoder := json.NewDecoder(file)
	err := decoder.Decode(userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func convertUserToFileContent(user *users.UserInfo) (string, error) {
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

func createToken(userID string, role users.Role, key []byte) (string, error) {
	jsonToken := paseto.JSONToken{
		Audience: "jotfly-user",
		Issuer: "jotfly-server",
		Jti: uuid.New().String(),
		Subject: userID,
		IssuedAt: time.Now(),
		Expiration: time.Now().AddDate(0, 0, 1),
		NotBefore: time.Now(),
	}
	jsonToken.Set("role", role.String())
	return paseto.NewV2().Encrypt(key, jsonToken, nil)
}

func checkGetToken(token string, key []byte) (userID string, role users.Role, err error) {
	jsonToken := &paseto.JSONToken{}
	err = paseto.NewV2().Decrypt(token, key, jsonToken, nil)
	if err != nil {
		return
	}
	if jsonToken.Expiration.Before(time.Now()) {
		err = errors.New("token expired")
		return
	}
	userID = jsonToken.Subject
	role = users.Role(users.Role_value[(jsonToken.Get("role"))])
	return
}

