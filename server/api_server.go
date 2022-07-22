package server

import (
	"github.com/choi-yh/example-golang/internal/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	router := s.router

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalf("Failed to Set Trusted Proxies")
	}

	//v1 := router.Group("/api/v1")
	//{
	//	v1.Any("/*Any")
	//}

	log.Printf("======= Start API Server on %s port =======", util.APIServerPort)
	if err := router.Run(":" + util.APIServerPort); err != nil {
		panic(err)
	}
}
