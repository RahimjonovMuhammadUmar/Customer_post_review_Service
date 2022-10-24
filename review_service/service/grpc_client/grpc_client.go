package grpcClient

import (
	"exam/review_service/config"
	pbp "exam/review_service/genproto/post"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcClientI ...
type GrpcClientI interface {
	Post() pbp.PostServiceClient
}

// GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	postService pbp.PostServiceClient
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
		cfg:         cfg,
		postService: pbp.NewPostServiceClient(connPost),
	}, nil
}

func (g *GrpcClient) Post() pbp.PostServiceClient {
    return g.postService
}