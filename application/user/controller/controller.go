package controller

import (
	"context"

	"github.com/choi-yh/example-golang/application/buf/gen/go/proto/user/v1"
	"github.com/choi-yh/example-golang/application/user/model"
	"github.com/choi-yh/example-golang/application/user/service"
	"github.com/pinpoint-apm/pinpoint-go-agent"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	user_v1.UnimplementedUserServiceServer
	svc service.Service
}

func RegisterServer(srv *grpc.Server) {
	user_v1.RegisterUserServiceServer(srv, &Server{
		svc: service.NewService(),
	})
}

func (s *Server) SignUp(ctx context.Context, request *user_v1.SignUpRequest) (*user_v1.SignUpResponse, error) {
	defer pinpoint.FromContext(ctx).NewSpanEvent("SignUp").EndSpanEvent()

	tracer := pinpoint.FromContext(ctx)
	pCtx := pinpoint.NewContext(context.Background(), tracer)

	user, err := s.svc.SignUp(pCtx, model.SignUpParam{
		Email:    request.Email,
		Password: request.Password,
		Name:     request.Name,
		Phone:    request.Phone,
	})
	if err != nil {
		return nil, err
	}

	return &user_v1.SignUpResponse{
		User: &user_v1.User{
			Id:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			Phone:     user.Phone,
			CreatedAt: timestamppb.New(user.CreatedAt),
		},
	}, nil
}

func (s *Server) Login(ctx context.Context, request *user_v1.LoginRequest) (*user_v1.LoginResponse, error) {

	return &user_v1.LoginResponse{}, nil
}
