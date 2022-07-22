package server

import (
	"log"
	"net"

	"github.com/choi-yh/example-golang/app"
	"github.com/choi-yh/example-golang/internal/util"
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

	lis, err := net.Listen("tcp", "localhost:"+util.GrpcServerPort)
	if err != nil {
		log.Fatalf("Failed to listen on %s port : %s", util.GrpcServerPort, err)
	}

	log.Printf("======= Start Grpc Server on %s Port =======", util.GrpcServerPort)
	if err = s.server.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve Grpc Server : %s", err)
	}
}
