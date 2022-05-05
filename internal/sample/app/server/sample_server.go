package server

import (
	"context"
	"fmt"
	samplepb "github.com/choi-yh/example-golang/protos/sample"
)

type SampleServer struct {
	samplepb.SampleServiceServer
}

func (s *SampleServer) Echo(ctx context.Context, request *samplepb.StringMessage) (*samplepb.StringMessage, error) {
	return &samplepb.StringMessage{Value: request.Value}, nil
}

func (s *SampleServer) SayHello(ctx context.Context, request *samplepb.HelloRequest) (*samplepb.HelloResponse, error) {
	helloMessage := fmt.Sprintf("Hello %s", request.Name)
	return &samplepb.HelloResponse{Message: helloMessage}, nil
}
