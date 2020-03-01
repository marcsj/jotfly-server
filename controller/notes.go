package controller

import (
	"github.com/marcsj/jotfly-server/notes"
	"github.com/marcsj/jotfly-server/repo"
)

type NotesController interface {
	GetDirectoryNotes(userID string, directory string) ([]string, error)
	GetNote(userID string, directory string, note string) (*notes.Note, error)
	CreateNote(userID string, directory string, note string) error
	UpdateNote(userID string, note *notes.Note) (*notes.Note, error)
	DeleteNote(userID string, directory string, note string) error
}

type notesController struct {
	usersRepo repo.UsersRepo
	notesRepo repo.NotesRepo
}

func NewNotesController(usersRepo repo.UsersRepo, notesRepo repo.NotesRepo) NotesController {
	return &notesController{
		usersRepo: usersRepo,
		notesRepo: notesRepo,
	}
}

func (c notesController) GetDirectoryNotes(userID string, directory string) ([]string, error) {
	return c.notesRepo.GetNotesInDirectory(userID, directory)
}

func (c notesController) GetNote(userID string, directory string, note string) (*notes.Note, error) {
	return c.notesRepo.GetNote(userID, directory, note)
}

func (c notesController) CreateNote(userID string, directory string, note string) error {
	user, err := c.usersRepo.GetUser(userID)
	if err != nil {
		return err
	}
	user.Directories = append(user.GetDirectories(), directory)
	_, err = c.usersRepo.UpdateUser(userID, user)
	if err != nil {
		return err
	}
	return c.notesRepo.CreateNote(userID, directory, note)
}

func (c notesController) UpdateNote(userID string, note *notes.Note) (*notes.Note, error) {
	return c.notesRepo.UpdateNote(userID, note)
}

func (c notesController) DeleteNote(userID string, directory string, note string) error {
	return c.notesRepo.DeleteNote(userID, directory, note)
}