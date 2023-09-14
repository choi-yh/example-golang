package server

import (
	user "github.com/choi-yh/example-golang/application/user/controller"
	"google.golang.org/grpc"
)

func Register(srv *grpc.Server) {
	user.RegisterServer(srv)
}
