package notes

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Repo interface {
	CreateNote(userID string, directory string, id string) error
	GetNote(userID string, directory string, id string) (*Note, error)
	UpdateNote(userID string, note *Note) (*Note, error)
	DeleteNote(userID string, directory string, id string) error
	GetNotesInDirectory(userID string, directory string) ([]string, error)
}

type fileRepo struct {
	path string
	mx sync.Mutex
}

func NewFileRepo(path string) *fileRepo {
	return &fileRepo{
		path: path,
	}
}

func (r fileRepo) CreateNote(userID string, directory string, id string) error {
	r.mx.Lock()
	defer r.mx.Unlock()
	err := os.MkdirAll(filepath.Join(r.path, userID, directory), 0744)
	if err != nil {
		return err
	}
	_, err = os.Create(filepath.Join(r.path, userID, directory, id))
	if err != nil {
		return err
	}
	return nil
}

func (r fileRepo) GetNote(userID string, directory string, id string) (*Note, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	file, err := os.Open(filepath.Join(r.path, userID, directory, id))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return convertFileToNote(userID, directory, id, file)
}

func (r fileRepo) UpdateNote(userID string, note *Note) (*Note, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	file, err := os.OpenFile(
		filepath.Join(r.path, userID, note.GetDirectory(), note.GetId()),
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0766,
		)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	noteContents, err := convertNoteToFileContent(note)
	if err != nil {
		return nil, err
	}
	_, err = file.WriteString(noteContents)
	if err != nil {
		return nil, err
	}

	file, err = os.Open(filepath.Join(r.path, userID, note.GetDirectory(), note.GetId()))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return convertFileToNote(userID, note.GetDirectory(), note.GetId(), file)
}

func (r fileRepo) DeleteNote(userID string, directory string, id string) error {
	r.mx.Lock()
	defer r.mx.Unlock()
	return os.Remove(filepath.Join(r.path, userID, directory, id))
}

func (r fileRepo) GetNotesInDirectory(userID string, directory string) ([]string, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	notes := make([]string, 0)
	err := filepath.Walk(filepath.Join(r.path, userID, directory),
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				name := strings.SplitAfter(path, "/")
				notes = append(notes, name[len(name)-1])
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return notes, nil
}

