package users

import (
	"io/ioutil"
	"os"
	"testing"
)

func setupFileRepo() (UserRepo, string, error) {
	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		return nil, dir, err
	}
	return NewFileRepo(dir), dir, nil
}

func TestFileRepo_CreateUser(t *testing.T) {
	repo, dir, err := setupFileRepo()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	err = repo.CreateUser("test@aol.com", []byte("test pass"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestFileRepo_UpdateNote(t *testing.T) {
	ownerID := "test@aol.com"
	pass := []byte("test pass")
	repo, dir, err := setupFileRepo()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	err = repo.CreateUser(ownerID, pass)
	if err != nil {
		t.Fatal(err)
	}
	user, err := repo.UpdateUser(ownerID, &UserInfo{
		Password: pass,
		Role: UserInfo_ADMIN,
	})
	if err != nil {
		t.Fatal(err)
	}
	if user.GetRole() != UserInfo_ADMIN {
		t.Fatal("user role is not admin")
	}
}

func TestFileRepo_DeleteNote(t *testing.T) {
	ownerID := "test@aol.com"
	pass := []byte("test pass")
	repo, dir, err := setupFileRepo()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	err = repo.CreateUser(ownerID, pass)
	if err != nil {
		t.Fatal(err)
	}
	err = repo.DeleteUser(ownerID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = repo.GetUser(ownerID)
	if err == nil {
		t.Fatal("user was not deleted")
	}
}