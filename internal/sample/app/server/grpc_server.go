package server

import (
	samplepb "github.com/choi-yh/example-golang/internal/sample/protos/sample"
	"google.golang.org/grpc"
	"log"
	"net"
)

const grpcServerPort = "9000"

type GrpcServer struct {
	server *grpc.Server
}

func NewGrpcServer() GrpcServer {
	return GrpcServer{
		server: grpc.NewServer(),
	}
}

func (s GrpcServer) Run() {
	samplepb.RegisterSampleServiceServer(s.server, &SampleServer{}) // Register SampleService on Grpc Server

	lis, err := net.Listen("tcp", "localhost:"+grpcServerPort)
	if err != nil {
		log.Fatalf("Failed to listen on %s port : %s", grpcServerPort, err)
	}

	log.Printf("======= Start Grpc Server on %s Port =======", grpcServerPort)
	if err = s.server.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve Grpc Server : %s", err)
	}
}
