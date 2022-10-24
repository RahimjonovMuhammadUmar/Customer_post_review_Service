package grpcClient

import (
	"exam/post_service/config"
	pbc "exam/post_service/genproto/customer"
	pbr "exam/post_service/genproto/review"

	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcClientI ...
type GrpcClientI interface {
	Customer() pbc.CustomerServiceClient
	Review() pbr.ReviewServiceClient
}

type GrpcClient struct {
	cfg              config.Config
	customerServices pbc.CustomerServiceClient
	reviewService    pbr.ReviewServiceClient
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {
	connCustomer, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CustomerServiceHost, cfg.CustomerServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error while connecting to connCustomer", err)
		return nil, err
	}
	connReview, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.ReviewServiceHost, cfg.ReviewServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error while connecting to connReview", err)
		return nil, err
	}

	return &GrpcClient{
		cfg:              cfg,
		customerServices: pbc.NewCustomerServiceClient(connCustomer),
		reviewService: pbr.NewReviewServiceClient(connReview),
	}, nil
}

func (g *GrpcClient) Customer() pbc.CustomerServiceClient {
	return g.customerServices
}

func (g *GrpcClient) Review() pbr.ReviewServiceClient {
	return g.reviewService
}