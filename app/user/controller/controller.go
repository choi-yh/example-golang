package controller

import (
	"context"

	"github.com/choi-yh/example-golang/app/user/service"
	userpb "github.com/choi-yh/example-golang/protos/user"
	"google.golang.org/grpc"
)

type Server struct {
	userpb.UnimplementedUserServiceServer
	svc service.Service
}

func RegisterServer(srv *grpc.Server) {
	userpb.RegisterUserServiceServer(srv, &Server{
		svc: service.NewService(),
	})
}

func (s *Server) SignUp(ctx context.Context, request *userpb.SignUpRequest) (*userpb.SignUpResponse, error) {

	return &userpb.SignUpResponse{}, nil
}

func (s *Server) Login(ctx context.Context, request *userpb.LoginRequest) (*userpb.LoginResponse, error) {

	return &userpb.LoginResponse{}, nil
}
