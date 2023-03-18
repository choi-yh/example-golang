package server

import (
	"fmt"
	"log"
	"net"

	"github.com/choi-yh/example-golang/internal/util/constants"

	"github.com/choi-yh/example-golang/app"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	server *grpc.Server
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{
		server: grpc.NewServer(),
	}
}

func (s *GrpcServer) Run() {
	app.Register(s.server)

	address := fmt.Sprintf("%s:%s", constants.HOST, constants.GrpcServerPort)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s port : %s", constants.GrpcServerPort, err)
	}

	log.Printf("======= Start Grpc Server on %s Port =======", constants.GrpcServerPort)
	if err = s.server.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve Grpc Server : %s", err)
	}
}
