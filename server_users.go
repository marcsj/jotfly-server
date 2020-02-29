package jotfly

import (
	"context"
	"github.com/marcsj/jotfly-server/users"
)

type usersServer struct {
}

func NewUsersServer() users.UsersServer {
	return &usersServer{}
}

func (s usersServer) GetDirectories(
	ctx context.Context,
	in *users.GetDirectoriesRequest) (*users.GetDirectoriesResponse, error) {
	return nil, nil
}

func (s usersServer) Login(
	ctx context.Context, in *users.User) (*users.Session, error) {
	return nil, nil
}