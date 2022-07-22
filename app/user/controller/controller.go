package controller

import (
	"context"

	userpb "github.com/choi-yh/example-golang/protos/user"
)

type Server struct {
	userpb.UnimplementedUserServiceServer
}

func (s *Server) SignUp(ctx context.Context, request *userpb.SignUpRequest) (*userpb.SignUpResponse, error) {

	return &userpb.SignUpResponse{}, nil
}

func (s *Server) Login(ctx context.Context, request *userpb.LoginRequest) (*userpb.LoginResponse, error) {

	return &userpb.LoginResponse{}, nil
}
