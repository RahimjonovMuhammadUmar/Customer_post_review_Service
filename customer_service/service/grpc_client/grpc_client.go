package grpcClient

import (
	"exam/customer_service/config"
	pbp "exam/customer_service/genproto/post"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcClientI ...
type GrpcClientI interface {
	Post() pbp.PostServiceClient
}

type GrpcClient struct {
	cfg          config.Config
	postServices pbp.PostServiceClient
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {
	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error while connecting to postService", err)
		return nil, err
	}

	return &GrpcClient{
		cfg:          cfg,
		postServices: pbp.NewPostServiceClient(connPost),
	}, nil
}

func (g *GrpcClient) Post() pbp.PostServiceClient {
	return g.postServices
}
