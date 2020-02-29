package notes

import (
	"io/ioutil"
	"os"
	"testing"
)

const testFileLocation = "./test"
func setupFileRepo() (NoteRepo, string, error) {
	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		return nil, dir, err
	}
	return NewFileRepo(dir), dir, nil
}

func TestFileRepo_CreateNote(t *testing.T) {
	repo, dir, err := setupFileRepo()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	err = repo.CreateNote("test@aol.com", "test", "test-note")
	if err != nil {
		t.Fatal(err)
	}
}