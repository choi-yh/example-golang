package server

import (
	"context"
	"fmt"

	"github.com/choi-yh/example-golang/application/sample/protos/sample"
)

type SampleServer struct {
	example_golang.example_golang
}

func (s *SampleServer) Echo(ctx context.Context, request *example_golang.example_golang) (*example_golang.StringMessage, error) {
	return &example_golang.StringMessage{Value: request.Value}, nil
}

func (s *SampleServer) SayHello(ctx context.Context, request *example_golang.HelloRequest) (*example_golang.HelloResponse, error) {
	helloMessage := fmt.Sprintf("Hello %s", request.Name)
	return &example_golang.HelloResponse{Message: helloMessage}, nil
}
