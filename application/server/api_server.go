package server

import (
	"context"
	"log"
	"net/http"

	userv1 "github.com/choi-yh/example-golang/application/buf/gen/go/proto/user/v1"
	"github.com/choi-yh/example-golang/application/internal/constants"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pinpoint-apm/pinpoint-go-agent"
	"github.com/pinpoint-apm/pinpoint-go-agent/plugin/gin"
	ppgrpc "github.com/pinpoint-apm/pinpoint-go-agent/plugin/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type APIServer struct {
	router *gin.Engine
}

func NewAPIServer() *APIServer {
	return &APIServer{
		router: gin.New(),
	}
}

func (s *APIServer) Run() {
	opts := []pinpoint.ConfigOption{
		pinpoint.WithAppName("APIServer"),
		pinpoint.WithAgentId("APIServerAgentID"),
		pinpoint.WithCollectorHost("localhost"),
	}
	cfg, _ := pinpoint.NewConfig(opts...)

	agent, err := pinpoint.NewAgent(cfg)
	if err != nil {
		log.Fatalf("APIServer pinpoint agent start fail: %v", err)
	}
	defer agent.Shutdown()

	ctx := context.Background()
	mux := runtime.NewServeMux()

	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),

		// pinpoint
		grpc.WithUnaryInterceptor(ppgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(ppgrpc.StreamClientInterceptor()),
	}

	// endpoint
	if err := userv1.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:"+constants.GrpcServerPort, options); err != nil {
		log.Fatalf("Failed Register Grpc Gateway on Grpc Server : %v", err)
	}

	router := s.router

	// middleware
	router.Use(
		ppgin.Middleware(),
	)

	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalf("Failed to Set Trusted Proxies")
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	v1 := router.Group("/v1")
	{
		v1.Any("/*Any", gin.WrapH(mux))
	}

	log.Printf("======= Start API Server on %s port =======", constants.APIServerPort)
	if err := router.Run(":" + constants.APIServerPort); err != nil {
		panic(err)
	}
}
