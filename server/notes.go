package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/marcsj/jotfly-server/notes"
)

type notesServer struct {
}

func NewNotesServer() notes.NotesServer {
	return &notesServer{}
}

func (s notesServer) GetDirectoryNotes(
	ctx context.Context,
	req *notes.GetDirectoryNotesRequest,
	) (*notes.GetDirectoryNotesResponse, error) {
	return nil, nil
}

func (s notesServer) GetNote(
	ctx context.Context,
	req *notes.GetNoteRequest,
	) (*notes.Note, error) {
	return nil, nil
}

func (s notesServer) CreateNote(
	ctx context.Context,
	req *notes.CreateRequest,
	) (*empty.Empty, error) {
	return nil, nil
}

func (s notesServer) UpdateNote(
	ctx context.Context,
	req *notes.UpdateRequest,
	) (*notes.Note, error) {
	return nil, nil
}

func (s notesServer) DeleteNote(
	ctx context.Context,
	req *notes.DeleteRequest,
	) (*empty.Empty, error) {
	return nil, nil
}