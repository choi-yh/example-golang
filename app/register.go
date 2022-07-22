package app

import (
	user "github.com/choi-yh/example-golang/app/user/controller"
	userpb "github.com/choi-yh/example-golang/protos/user"
	"google.golang.org/grpc"
)

func Register(srv *grpc.Server) {
	userpb.RegisterUserServiceServer(srv, &user.Server{})
}
