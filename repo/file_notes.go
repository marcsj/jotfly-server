package repo

import (
	"github.com/marcsj/jotfly-server/notes"
	"github.com/marcsj/jotfly-server/util"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type NotesRepo interface {
	CreateNote(userID string, directory string, id string) error
	GetNote(userID string, directory string, id string) (*notes.Note, error)
	UpdateNote(userID string, note *notes.Note) (*notes.Note, error)
	DeleteNote(userID string, directory string, id string) error
	GetNotesInDirectory(userID string, directory string) ([]string, error)
}

type fileNotesRepo struct {
	path string
	mx sync.Mutex
}

func NewFileNotesRepo(path string) *fileNotesRepo {
	return &fileNotesRepo{
		path: path,
	}
}

func (r fileNotesRepo) CreateNote(userID string, directory string, id string) error {
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

func (r fileNotesRepo) GetNote(userID string, directory string, id string) (*notes.Note, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	file, err := os.Open(filepath.Join(r.path, userID, directory, id))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return util.ConvertFileToNote(userID, directory, id, file)
}

func (r fileNotesRepo) UpdateNote(userID string, note *notes.Note) (*notes.Note, error) {
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

	noteContents, err := util.ConvertNoteToFileContent(note)
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
	return util.ConvertFileToNote(userID, note.GetDirectory(), note.GetId(), file)
}

func (r fileNotesRepo) DeleteNote(userID string, directory string, id string) error {
	r.mx.Lock()
	defer r.mx.Unlock()
	return os.Remove(filepath.Join(r.path, userID, directory, id))
}

func (r fileNotesRepo) GetNotesInDirectory(userID string, directory string) ([]string, error) {
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

