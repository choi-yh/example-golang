package main

import (
	"github.com/choi-yh/example-golang/sample/app"
	"github.com/choi-yh/example-golang/sample/app/server"
)

func main() {
	grpcGateway := app.NewGrpcGateway()
	go grpcGateway.Run()

	grpcServer := server.NewGrpcServer()
	grpcServer.Run()
}
