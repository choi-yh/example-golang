package main

import "github.com/choi-yh/example-golang/sample/app"

func main() {
	grpcServer := app.NewGrpcServer()
	grpcServer.Run()
}
