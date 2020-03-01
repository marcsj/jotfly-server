package repo

import (
	"github.com/marcsj/jotfly-server/notes"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func setupFileNotesRepo() (NotesRepo, string, error) {
	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		return nil, dir, err
	}
	return NewFileNotesRepo(dir), dir, nil
}

func TestFileNotesRepo_CreateNote(t *testing.T) {
	repo, dir, err := setupFileNotesRepo()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	err = repo.CreateNote("test@aol.com", "test", "test-note")
	if err != nil {
		t.Fatal(err)
	}
}

func TestFileNotesRepo_UpdateNote(t *testing.T) {
	ownerID := "test@aol.com"
	directory := "test"
	id := "test-note"
	content := "This is a quality note!"
	repo, dir, err := setupFileNotesRepo()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	err = repo.CreateNote(ownerID, directory, id)
	if err != nil {
		t.Fatal(err)
	}
	newNote, err := repo.UpdateNote(ownerID, &notes.Note{
		OwnerId: ownerID,
		Directory: directory,
		Id: id,
		Content: content,
	})
	if err != nil {
		t.Fatal(err)
	}
	if newNote.GetContent() != content {
		t.Fatal("updated note does not have the same content")
	}
}

func TestFileNotesRepo_DeleteNote(t *testing.T) {
	ownerID := "test@aol.com"
	directory := "test"
	id := "test-note"
	repo, dir, err := setupFileNotesRepo()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	err = repo.CreateNote(ownerID, directory, id)
	if err != nil {
		t.Fatal(err)
	}
	err = repo.DeleteNote(ownerID, directory, id)
	if err != nil {
		t.Fatal(err)
	}
	_, err = repo.GetNote(ownerID, directory, id)
	if err == nil {
		t.Fatal("note was not deleted")
	}
}

func TestFileNotesRepo_GetNotesInDirectory(t *testing.T) {
	ownerID := "test@aol.com"
	directory := "test"
	ids := []string{"test1", "test2", "test3"}
	repo, dir, err := setupFileNotesRepo()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	for _, id := range ids {
		err = repo.CreateNote(ownerID, directory, id)
		if err != nil {
			t.Fatal(err)
		}
	}
	foundIDs, err := repo.GetNotesInDirectory(ownerID, directory)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(ids, foundIDs) {
		t.Fatal("list is not the same")
	}
}
