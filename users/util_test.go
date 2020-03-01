package users

import "testing"

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
