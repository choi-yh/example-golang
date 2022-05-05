package app

import (
	"context"
	samplepb "github.com/choi-yh/example-golang/internal/sample/protos/sample"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

const (
	grpcGatewayPort = "8080"
	grpcServerPort  = "9000"
)

type GrpcGateway struct {
	server *gin.Engine
}

func NewGrpcGateway() GrpcGateway {
	return GrpcGateway{
		server: gin.New(),
	}
}

func (s GrpcGateway) Run() {
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
		log.Fatalf("Failed Register Grpc Gateway on Grpc Server : %v", err)
	}

	r := s.server
	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalf("Failed to Set Trusted Proxies")
	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	v1 := r.Group("/api/v1")
	{
		v1.Any("/*Any", gin.WrapH(mux))
	}

	log.Printf("======= Start Grpc Gateway on %s port =======", grpcGatewayPort)
	if err := r.Run(":" + grpcGatewayPort); err != nil {
		log.Fatalf("Failed to Run Grpc Gateway : %s", err)
	}
}
