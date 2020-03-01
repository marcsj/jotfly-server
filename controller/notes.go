package controller

import (
	"github.com/marcsj/jotfly-server/repo"
)

type NotesController interface {
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
