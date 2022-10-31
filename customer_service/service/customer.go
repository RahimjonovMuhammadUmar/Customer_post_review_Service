package service

import (
	"context"
	"database/sql"
	pbc "exam/customer_service/genproto/customer"
	"fmt"

	pbp "exam/customer_service/genproto/post"
	l "exam/customer_service/pkg/logger"
	grpcClient "exam/customer_service/service/grpc_client"
	"exam/customer_service/storage"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CustomerService ...
type CustomerService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.GrpcClientI
}

// NewCustomerService ...
func NewCustomerService(db *sqlx.DB, log l.Logger, client grpcClient.GrpcClientI) *CustomerService {
	return &CustomerService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (c *CustomerService) CreateCustomer(ctx context.Context, req *pbc.CustomerRequest) (*pbc.CustomerWithoutPost, error) {
	createdCustomer, err := c.storage.Customer().CreateCustomer(req)
	if err != nil {
		c.logger.Error("error -> createdCustomer, err := c.storage.Customer().CreateCustomer(req)", l.Any("error CreatingCustomer grpc_client/customer.go", err))
		return &pbc.CustomerWithoutPost{}, err
	}
	return createdCustomer, nil
}
func (c *CustomerService) UpdateCustomer(ctx context.Context, req *pbc.CustomerWithoutPost) (*pbc.CustomerWithoutPost, error) {
	_, err := c.storage.Customer().UpdateCustomer(req)
	if err != nil {
		c.logger.Error("error -> _, err := c.storage.Customer().UpdateCustomer(req)", l.Any("error UpdatingCustomer grpc_client/customer.go", err))
		return &pbc.CustomerWithoutPost{}, nil
	}
	return req, nil
}
func (c *CustomerService) CheckIfCustomerExists(ctx context.Context, req *pbc.CustomerId) (*pbc.Exists, error) {
	exist, err := c.storage.Customer().CheckIfCustomerExists(req.Id)
	if err != nil {
		c.logger.Error("error -> exist, err := c.storage.Customer().CheckIfCustomerExists(req.Id)", l.Any("error Checking for customer existence grpc_client/customer.go", err))
		return &pbc.Exists{}, err
	}
	if exist.Exists {
		return &pbc.Exists{Exists: true}, nil
	}
	return &pbc.Exists{Exists: false}, nil
}
func (c *CustomerService) GetCustomer(ctx context.Context, req *pbc.CustomerId) (*pbc.Customer, error) {
	customerData, err := c.storage.Customer().GetCustomer(req.Id)
	if err == sql.ErrNoRows {
		c.logger.Info("No such customer")
		return &pbc.Customer{}, status.Error(codes.NotFound, "There is no such customer")

	}
	if err != nil {
		c.logger.Error("error -> customerData, err := c.storage.Customer().GetCustomer(req.Id)", l.Any("error getting customer by id grpc_client/customer.go", err))
		return &pbc.Customer{}, err
	}

	Posts, err := c.client.Post().GetPostsOfCustomer(ctx, &pbp.Id{Id: req.Id})
	if err != nil {
		c.logger.Error("error while sending req to GetPostsOfCustomer", l.Any("", err))
		return nil, err
	}

	for _, post := range Posts.Posts {
		CustomerPost := &pbc.Post{
			Id:          post.Id,
			Name:        post.Name,
			Description: post.Description,
		}
		for _, media := range post.Medias {
			Customer_media := &pbc.Media{
				Id:   media.Id,
				Name: media.Name,
				Link: media.Link,
				Type: media.Type,
			}
			CustomerPost.Medias = append(CustomerPost.Medias, Customer_media)
		}
		for _, review := range post.Reviews {
			cust_review := &pbc.Review{
				Id:         review.Id,
				Name:       review.Name,
				Review:     review.Review,
				CustomerId: review.CustomerId,
			}
			CustomerPost.Reviews = append(CustomerPost.Reviews, cust_review)
		}

		customerData.Posts = append(customerData.Posts, CustomerPost)
	}

	return customerData, nil
}
func (c *CustomerService) DeleteCustomer(ctx context.Context, req *pbc.CustomerId) (*pbc.CustomerDeleted, error) {
	deleted, err := c.storage.Customer().DeleteCustomer(req.Id)
	if err != nil {
		c.logger.Error("error while deleting service/customer.go", l.Any("", err))
		return &pbc.CustomerDeleted{}, err
	}
	if !deleted.CustomerDeleted {
		return &pbc.CustomerDeleted{
			CustomerDeleted: false,
		}, nil
	}

	_, err = c.client.Post().DeletePostByCustomerId(ctx, &pbp.Id{
		Id: req.Id,
	})
	if err != nil {
		c.logger.Error("error while sending id to delete posts", l.Any("", err))
	}

	return deleted, nil
}
func (c *CustomerService) CheckField(ctx context.Context, req *pbc.FieldCheck) (*pbc.Exists, error) {
	exist, err := c.storage.Customer().CheckField(req.Field, req.EmailOrUsername)
	if err != nil {
		c.logger.Error("error while sending field to db to check", l.Any("", err))
		return &pbc.Exists{}, err
	}
	if !exist.Exists {
		return &pbc.Exists{
			Exists: false,
		}, nil
	}
	return &pbc.Exists{
		Exists: true,
	}, nil
}

func (c *CustomerService) SearchCustomer(ctx context.Context, req *pbc.InfoForSearch) (*pbc.PossibleCustomers, error) {
	posts, err := c.storage.Customer().SearchCustomer(req.Field, req.Value, req.OrderBy, req.AscOrDesc, req.Limit, req.Page)
	if err != nil {
		c.logger.Error("error while sending to db to seacrh", l.Error(err))
		return &pbc.PossibleCustomers{}, err
	}
	fmt.Println(err)

	return posts, nil
}
