package app

import (
	"context"
	samplepb "github.com/choi-yh/example-golang/protos/sample"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

const (
	grpcServerPort = "9000"
)

type GrpcServer struct {
	server *grpc.Server
}

func NewGrpcServer() GrpcServer {
	return GrpcServer{
		server: grpc.NewServer(),
	}
}

func (s GrpcServer) Run() {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := samplepb.RegisterSampleServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:"+grpcServerPort,
		options,
	); err != nil {
		log.Fatalf("Failed Register Sample Service on Grpc Server : %v", err)
	}

	lis, err := net.Listen("tcp", "localhost:"+grpcServerPort)
	if err != nil {
		log.Printf("Failed to Listen on %s port", grpcServerPort)
	}

	log.Printf("======== Start Grpc Server on %s port =======", grpcServerPort)
	if err = s.server.Serve(lis); err != nil {
		log.Printf("Failed to Serve Grpc Server : %v", err)
	}
}
