package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/choi-yh/example-golang/internal/util/constants"
	oauthpb "github.com/choi-yh/example-golang/protos/oauth"

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
	bctx := context.Background()
	ctx, cancel := context.WithCancel(bctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := registerEndpoint(ctx, mux, opts); err != nil {
		log.Fatalf("Failed Register Endpoint : %v", err)
	}

	log.Printf("======= Start API Server on %s port =======", constants.APIServerPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", constants.APIServerPort), mux); err != nil {
		log.Fatalf("Failed Register API Gateway on API Server : %v", err)
	}
}

func registerEndpoint(ctx context.Context, mux *runtime.ServeMux, opts []grpc.DialOption) error {
	endpoint := fmt.Sprintf("%s:%s", constants.HOST, constants.GrpcServerPort)

	if err := userpb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		//log.Fatalf("Failed Register Grpc Gateway on Grpc User Server : %v", err)
		return err
	}

	if err := oauthpb.RegisterOAuthServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		//log.Fatalf("Failed Register Grpc Gateway on Grpc OAuth Server : %v", err)
		return err
	}

	return nil
}
