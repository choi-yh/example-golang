package main

import (
	"github.com/choi-yh/example-golang/server"
)

func main() {
	apiServer := server.NewAPIServer()
	go apiServer.Run()

	grpcServer := server.NewGrpcServer()
	grpcServer.Run()
}
