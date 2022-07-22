package server

import (
	"context"
	"log"
	"net/http"

	"github.com/choi-yh/example-golang/internal/util"
	userpb "github.com/choi-yh/example-golang/protos/user"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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
	ctx := context.Background()
	mux := runtime.NewServeMux()

	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := userpb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:"+util.GrpcServerPort, options); err != nil {
		log.Fatalf("Failed Register Grpc Gateway on Grpc Server : %v", err)
	}

	router := s.router

	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalf("Failed to Set Trusted Proxies")
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	v1 := router.Group("/api/v1")
	{
		v1.Any("/*Any", gin.WrapH(mux))
	}

	log.Printf("======= Start API Server on %s port =======", util.APIServerPort)
	if err := router.Run(":" + util.APIServerPort); err != nil {
		panic(err)
	}
}
