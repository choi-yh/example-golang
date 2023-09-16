package server

import (
	"log"
	"net"

	"github.com/choi-yh/example-golang/application/internal/constants"
	"github.com/pinpoint-apm/pinpoint-go-agent"
	ppgrpc "github.com/pinpoint-apm/pinpoint-go-agent/plugin/grpc"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	server *grpc.Server
}

func NewGrpcServer() *GrpcServer {
	grpcServer := grpc.NewServer(
		// pinpoint
		grpc.UnaryInterceptor(ppgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(ppgrpc.StreamServerInterceptor()),
	)

	return &GrpcServer{
		server: grpcServer,
	}
}

func (s *GrpcServer) Run() {
	opts := []pinpoint.ConfigOption{
		pinpoint.WithAppName("gRPCServer"),
		pinpoint.WithAgentId("gRPCServerAgentID"),
		pinpoint.WithCollectorHost("localhost"),
	}
	cfg, _ := pinpoint.NewConfig(opts...)

	agent, err := pinpoint.NewAgent(cfg)
	if err != nil {
		log.Fatalf("gRPCServer pinpoint agent start fail: %v", err)
	}
	defer agent.Shutdown()

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
