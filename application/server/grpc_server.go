package server

import (
	"log"
	"net"

	"github.com/choi-yh/example-golang/application/internal/constants"
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
	Register(s.server)

	lis, err := net.Listen("tcp", "localhost:"+constants.GrpcServerPort)
	if err != nil {
		log.Fatalf("Failed to listen on %s port : %s", constants.GrpcServerPort, err)
	}

	log.Printf("======= Start Grpc Server on %s Port =======", constants.GrpcServerPort)
	if err = s.server.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve Grpc Server : %s", err)
	}
}
