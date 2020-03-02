package server

import (
	"context"
	"github.com/marcsj/jotfly-server/repo"
	"github.com/marcsj/jotfly-server/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)


type AuthMiddleware interface {
	UnaryAuthOption() grpc.ServerOption
}

type authMiddleware struct {
	salt []byte
	usersRepo repo.UsersRepo
}

func NewAuthMiddleware(salt []byte, usersRepo repo.UsersRepo) AuthMiddleware {
	return &authMiddleware{salt: salt}
}

func (m authMiddleware) serverAuthInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (hand interface{}, err error) {

	// ignore on login
	if info.FullMethod == "/users.Users/Login" {
		return handler(ctx, req), nil
	}

	ctx, err = m.authCheck(ctx)
	if err != nil {
		return nil, err
	}


	return handler(ctx, req)
}

func (m authMiddleware) authCheck(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.InvalidArgument, "metadata failure")
	}
	authToken := md.Get("authorization")
	if len(authToken) < 1 {
		return ctx, status.Errorf(codes.Unauthenticated, "authorization token not set")
	}
	token := authToken[0]
	id, role, err := util.CheckGetToken(token, m.salt)
	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "authentication failure")
	}
	_, err = m.usersRepo.GetUser(id)
	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "authentication failure")
	}

	// set our user info for further calls
	md.Set("id", id)
	md.Set("role", role.String())
	return metadata.NewIncomingContext(ctx, md), nil
}

func (m authMiddleware) UnaryAuthOption() grpc.ServerOption {
	return grpc.UnaryInterceptor(
		m.serverAuthInterceptor)
}