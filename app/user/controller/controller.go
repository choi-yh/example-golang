package controller

import (
	"context"

	"github.com/choi-yh/example-golang/app/user/model"
	"github.com/choi-yh/example-golang/app/user/service"
	userpb "github.com/choi-yh/example-golang/protos/user"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	user, err := s.svc.SignUp(ctx, model.SignUpParam{
		Email:    request.Email,
		Password: request.Password,
		Name:     request.Name,
		Phone:    request.Phone,
	})
	if err != nil {
		return nil, err
	}

	return &userpb.SignUpResponse{
		User: &userpb.User{
			Id:        int32(user.ID),
			Email:     user.Email,
			Name:      user.Name,
			Phone:     user.Phone,
			CreatedAt: timestamppb.New(*user.CreatedAt),
			UpdatedAt: timestamppb.New(*user.UpdatedAt),
		},
	}, nil
}

func (s *Server) Login(ctx context.Context, request *userpb.LoginRequest) (*userpb.LoginResponse, error) {

	return &userpb.LoginResponse{}, nil
}
