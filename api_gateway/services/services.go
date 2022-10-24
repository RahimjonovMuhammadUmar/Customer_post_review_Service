package services

import (
	"exam/api_gateway/config"
	pbc "exam/api_gateway/genproto/customer"
	pbp "exam/api_gateway/genproto/post"
	pbr "exam/api_gateway/genproto/review"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	CustomerService() pbc.CustomerServiceClient
	ReviewService() pbr.ReviewServiceClient
	PostService() pbp.PostServiceClient
}

type serviceManager struct {
	customerService pbc.CustomerServiceClient
	reviewService   pbr.ReviewServiceClient
	postService     pbp.PostServiceClient
}

func (s *serviceManager) CustomerService() pbc.CustomerServiceClient {
	return s.customerService
}
func (s *serviceManager) ReviewService() pbr.ReviewServiceClient {
	return s.reviewService
}
func (s *serviceManager) PostService() pbp.PostServiceClient {
	return s.postService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connCustomer, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CustomerServiceHost, conf.CustomerServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error while connecting to Customer api_gateway/services/customer", err)
		return nil, err
	}

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error while connecting to Post api_gateway/services/post", err)
		return nil, err
	}

	connReview, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.ReviewServiceHost, conf.ReviewServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error while connecting to Review api_gateway/services/post", err)
		return nil, err
	}

	serviceManager := &serviceManager{
		customerService: pbc.NewCustomerServiceClient(connCustomer),
		postService:     pbp.NewPostServiceClient(connPost),
		reviewService:   pbr.NewReviewServiceClient(connReview),
	}

	return serviceManager, nil
}
