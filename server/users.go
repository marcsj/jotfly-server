package server

import (
	"context"
	"errors"
	"github.com/marcsj/jotfly-server/controller"
	"github.com/marcsj/jotfly-server/users"
	"google.golang.org/grpc/metadata"
)

type usersServer struct {
	controller controller.UsersController
}

func NewUsersServer(controller controller.UsersController) users.UsersServer {
	return &usersServer{
		controller: controller,
	}
}

func (s usersServer) GetDirectories(
	ctx context.Context,
	in *users.GetDirectoriesRequest) (*users.GetDirectoriesResponse, error) {
	id, err := getId(ctx)
	if err != nil {
		return nil, err
	}
	directories, err := s.controller.GetDirectories(id)
	if err != nil {
		return nil, err
	}
	return &users.GetDirectoriesResponse{Directories: directories}, nil
}

func (s usersServer) Login(
	ctx context.Context, in *users.User) (*users.Session, error) {
	token, err := s.controller.Login(in.GetId(), in.GetPassword())
	if err != nil {
		return nil, err
	}
	return &users.Session{Token: token}, nil
}

func getId(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no id in context")
	}
	id := md.Get("id")
	if len(id) < 1 {
		return "", errors.New("no id in context")
	}
	return id[0], nil
}