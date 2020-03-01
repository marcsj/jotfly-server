package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/marcsj/jotfly-server/controller"
	"github.com/marcsj/jotfly-server/notes"
)

type notesServer struct {
	controller controller.NotesController
}

func NewNotesServer(controller controller.NotesController) notes.NotesServer {
	return &notesServer{
		controller: controller,
	}
}

func (s notesServer) GetDirectoryNotes(
	ctx context.Context,
	req *notes.GetDirectoryNotesRequest,
	) (*notes.GetDirectoryNotesResponse, error) {
	id, err := getId(ctx)
	if err != nil {
		return nil, err
	}
	listNotes, err := s.controller.GetDirectoryNotes(id, req.GetDirectory())
	if err != nil {
		return nil, err
	}
	return &notes.GetDirectoryNotesResponse{
		Notes: listNotes,
	}, nil
}

func (s notesServer) GetNote(
	ctx context.Context,
	req *notes.GetNoteRequest,
	) (*notes.Note, error) {
	id, err := getId(ctx)
	if err != nil {
		return nil, err
	}
	return s.controller.GetNote(id, req.GetDirectory(), req.GetId())
}

func (s notesServer) CreateNote(
	ctx context.Context,
	req *notes.CreateRequest,
	) (*empty.Empty, error) {
	id, err := getId(ctx)
	if err != nil {
		return nil, err
	}
	return nil, s.controller.CreateNote(id, req.GetDirectory(), req.GetId())
}

func (s notesServer) UpdateNote(
	ctx context.Context,
	req *notes.UpdateRequest,
	) (*notes.Note, error) {
	id, err := getId(ctx)
	if err != nil {
		return nil, err
	}
	return s.controller.UpdateNote(id, req)
}

func (s notesServer) DeleteNote(
	ctx context.Context,
	req *notes.DeleteRequest,
	) (*empty.Empty, error) {
	id, err := getId(ctx)
	if err != nil {
		return nil, err
	}
	return nil, s.controller.DeleteNote(id, req.GetDirectory(), req.GetId())
}