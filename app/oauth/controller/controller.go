package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/choi-yh/example-golang/app/user/service"
	"github.com/choi-yh/example-golang/internal/config"
	oauthpb "github.com/choi-yh/example-golang/protos/oauth"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	oauthpb.UnimplementedOAuthServiceServer
	svc service.Service
}

func RegisterOAuthServer(srv *grpc.Server) {
	oauthpb.RegisterOAuthServiceServer(srv, &Server{
		svc: service.NewService(),
	})
}

func (s *Server) GoogleLogin(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	googleOAuthConfig := config.GetGoogleOAuthConfig()
	_ = googleOAuthConfig.AuthCodeURL(config.State)
	//TODO: redirect

	return &emptypb.Empty{}, nil
}

func (s *Server) GoogleCallback(ctx context.Context, in *oauthpb.GoogleCallbackRequest) (*oauthpb.GoogleCallbackResponse, error) {
	//TODO: state 검증

	//TODO: token 교환
	googleOAuthConfig := config.GetGoogleOAuthConfig()
	token, err := googleOAuthConfig.Exchange(ctx, in.Code)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	if err != nil {
		fmt.Printf("failed to create user info request: %s\n", err.Error())
		//http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("failed to fetch user info: %s\n", err.Error())
		//http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil, err
	}
	defer response.Body.Close()

	res := &oauthpb.GoogleCallbackResponse{}

	return res, nil
}
