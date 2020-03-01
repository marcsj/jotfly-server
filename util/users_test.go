package util

import (
	"github.com/marcsj/jotfly-server/users"
	"testing"
)

func Test_Password(t *testing.T) {
	salt, err := generateRandomBytes(16)
	if err != nil {
		t.Fatal(err)
	}
	pass := generatePassword("test pass123", salt)
	if err != nil {
		t.Fatal(err)
	}
	err = checkPassword("test pass123", pass, salt)
	if err != nil {
		t.Fatal(err)
	}
	err = checkPassword("test hack123", pass, salt)
	if err == nil {
		t.Fatal("passwords are not actually the same, hacker infiltration")
	}
}

func Test_CreateToken(t *testing.T) {
	id := "test@aol.com"
	key := []byte("YELLOW SUBMARINE, BLACK WIZARDRY")
	token, err := createToken(id, users.Role_ADMIN, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
	userID, role, err := checkGetToken(token, key)
	if err != nil {
		t.Fatal(err)
	}
	if userID != id || role != users.Role_ADMIN {
		t.Fatal("payload came back wrong")
	}
}