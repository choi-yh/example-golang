package app

import (
	oauth "github.com/choi-yh/example-golang/app/oauth/controller"
	user "github.com/choi-yh/example-golang/app/user/controller"
	"google.golang.org/grpc"
)

func Register(srv *grpc.Server) {
	user.RegisterServer(srv)
	oauth.RegisterOAuthServer(srv)
}
