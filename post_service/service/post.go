package service

import (
	"context"
	"database/sql"
	pbc "exam/post_service/genproto/customer"
	pbp "exam/post_service/genproto/post"
	pbr "exam/post_service/genproto/review"
	l "exam/post_service/pkg/logger"
	grpcClient "exam/post_service/service/grpc_client"
	"exam/post_service/storage"
	"fmt"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PostService ...
type PostService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.GrpcClientI
}

// NewPostService ...
func NewPostService(db *sqlx.DB, log l.Logger, client grpcClient.GrpcClientI) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (p *PostService) CreatePost(ctx context.Context, req *pbp.PostRequest) (*pbp.PostWithoutReview, error) {
	exists, err := p.client.Customer().CheckIfCustomerExists(ctx, &pbc.CustomerId{
		Id: req.CustomerId,
	})
	if err != nil {
		p.logger.Error("error -> exists, err := p.client.Customer().CheckIfCustomerExists(ctx, &pbc.CustomerId{", l.Any("error checking customer by id post/service/grpc_client/customer.go", err))
		return &pbp.PostWithoutReview{}, err
	}
	if !exists.Exists {
		p.logger.Info("There is no such customer")

		return &pbp.PostWithoutReview{}, status.Error(codes.NotFound, "There is no such customer")
	}

	newPost, err := p.storage.Post().CreatePost(req)
	if err != nil {
		p.logger.Error("error -> newPost, err :=", l.Any("error creating post post/service/grpc_client/customer.go", err))
		return &pbp.PostWithoutReview{}, err
	}

	return newPost, nil
}
func (p *PostService) GetPostWithCustomerInfo(ctx context.Context, req *pbp.Ids) (*pbp.PostWithCustomerInfo, error) {
	customer, err := p.client.Customer().GetCustomer(ctx, &pbc.CustomerId{
		Id: req.CustomerId,
	})
	if err != nil {
		p.logger.Error("error -> customer, err := p.client.Customer().IfCustomerExists(ctx, &pbc.CustomerId{", l.Any("", err))
		return &pbp.PostWithCustomerInfo{}, err
	}
	post, err := p.storage.Post().GetPostWithCustomerInfo(req.Id)
	if err != nil && err != sql.ErrNoRows {
		p.logger.Error("error -> post, err := p.storage.Post().GetPost(req)", l.Any("error getting post by id post/service/post.go", err))
		return &pbp.PostWithCustomerInfo{}, err
	}

	post.Customer = &pbp.Customer{
		Id:          customer.Id,
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		Bio:         customer.Bio,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
	}
	for _, address := range customer.Addresses {
		address_ := &pbp.Address{
			Id:          address.Id,
			Street:      address.Street,
			HouseNumber: address.HouseNumber,
		}
		post.Customer.Addresses = append(post.Customer.Addresses, address_)
	}
	reviews, err := p.client.Review().GetReviews(ctx, &pbr.ReviewPostId{
		PostId: post.Id,
	})
	if err != nil {
		p.logger.Error("error -> reviews, err := p.client.Review().GetReviews(ctx, &pbr.ReviewPostId", l.Any("error getting reviews by post_id post/service/post.go", err))
		return nil, err
	}

	for _, review := range reviews.Reviews {
		post.Reviews = append(post.Reviews, &pbp.Review{
			Id:          review.Id,
			Name:        review.Name,
			Review:      review.Review,
			Description: review.Description,
			PostId:      review.PostId,
		})
	}
	return post, nil
}
func (p *PostService) UpdatePost(ctx context.Context, req *pbp.PostWithoutReview) (*pbp.PostWithoutReview, error) {
	updatedPost, err := p.storage.Post().UpdatePost(req)
	if err != nil {
		p.logger.Error("error -> exists, err := p.client.Customer().IfCustomerExists(ctx, &pbc.CustomerId{", l.Any("error checking customer by id post/service/grpc_client/customer.go", err))
		return &pbp.PostWithoutReview{}, err
	}
	return updatedPost, nil
}
func (p *PostService) DoesPostExist(ctx context.Context, req *pbp.Id) (*pbp.Exist, error) {
	// exist, err := p.client.Customer().CheckIfCustomerExists(ctx, &pbc.CustomerId{
	// 	Id: req.Id,
	// })
	// if err != nil {
	// 	p.logger.Error("exist, err := p.client.Customer().CheckIfCustomerExists(ctx, &pbc.CustomerId{", l.Any("error while checking for customer by id", err))
	// 	return &pbp.Exist{
	// 		Exists: false,
	// 	}, err
	// }
	// if !exist.Exists {
	// 	p.logger.Info("such customer does not exist")
	// 	return &pbp.Exist{
	// 		Exists: false,
	// 	}, nil
	// }
	exists, err := p.storage.Post().DoesPostExist(req.Id)
	if err != nil {
		p.logger.Error("Error exists, err = p.storage.Post().DoesPostExist(req.Id)")
		return nil, err
	}
	if !exists.Exists {
		p.logger.Info("such post does not exist")
		return &pbp.Exist{}, status.Error(codes.NotFound, "There is no such post")

	}

	return &pbp.Exist{
		Exists: true,
	}, nil
}
func (p *PostService) DeletePost(ctx context.Context, req *pbp.Id) (*pbp.IsDeleted, error) {
	deleted, err := p.storage.Post().DeletePost(req.Id)
	if err != nil {
		p.logger.Error("error while deleting grpc_client/post.go", l.Any("", err))
		return &pbp.IsDeleted{
			PostDeleted: false,
		}, err
	}
	if !deleted.PostDeleted {
		return &pbp.IsDeleted{
			PostDeleted: false,
		}, nil
	}
	_, err = p.client.Review().DeletePostsReviews(ctx, &pbr.ReviewPostId{
		PostId: req.Id,
	})
	if err != nil {
		p.logger.Error("error while sending id to reviews of post", l.Any("", err))
		return &pbp.IsDeleted{
			PostDeleted: false,
		}, err
	}
	return &pbp.IsDeleted{
		PostDeleted: true,
	}, nil
}

func (p *PostService) GetAllPostsWithCustomer(ctx context.Context, req *pbp.Empty) (*pbp.AllPosts, error) {
	allPosts, err := p.storage.Post().GetAllPostsWithCustomer(req)
	if err != nil {
		p.logger.Error("error while sending request to db GetAllPosts", l.Any("", err))
		return &pbp.AllPosts{}, err
	}
	fmt.Println(allPosts)
	for _, post := range allPosts.Posts {
		fmt.Println(post)
		customer, err := p.client.Customer().GetCustomer(ctx, &pbc.CustomerId{
			Id: post.CustomerId,
		})
		if err != nil {
			p.logger.Error("error while sending req to get customer", l.Any("", err))
			return &pbp.AllPosts{}, err
		}
		fmt.Println(customer)
		post.Customer = &pbp.Customer{
			Id:          customer.Id,
			FirstName:   customer.FirstName,
			LastName:    customer.LastName,
			Bio:         customer.Bio,
			Email:       customer.Email,
			PhoneNumber: customer.PhoneNumber,
		}
		for _, address := range customer.Addresses {
			fmt.Println(address)

			post.Customer.Addresses = append(post.Customer.Addresses, &pbp.Address{
				Id:          address.Id,
				Street:      address.Street,
				HouseNumber: address.HouseNumber,
			})
		}
	}

	fmt.Println(allPosts)
	return allPosts, nil
}

func (p *PostService) GetPostsOfCustomer(ctx context.Context, req *pbp.Id) (*pbp.Posts, error) {
	exists, err := p.client.Customer().CheckIfCustomerExists(ctx, &pbc.CustomerId{
		Id: req.Id,
	})
	if err != nil {
		p.logger.Error("error -> exists, err := p.client.Customer().CheckIfCustomerExists(ctx, &pbc.CustomerId{", l.Any("error checking customer by id post/service/grpc_client/customer.go", err))
		return &pbp.Posts{}, err
	}
	if !exists.Exists {
		p.logger.Info("There is no such customer")

		return &pbp.Posts{}, status.Error(codes.NotFound, "There is no such customer")
	}

	posts, err := p.storage.Post().GetPostsOfCustomer(req)
	if err != nil {
		p.logger.Error("error while sending to db GetPostsOfCustomer", l.Any("", err))
		return &pbp.Posts{}, err
	}
	for _, post := range posts.Posts {
		reviews, err := p.client.Review().GetReviews(ctx, &pbr.ReviewPostId{
			PostId: post.Id,
		})
		if err != nil {
			p.logger.Error("error while sending post_id to reviews", l.Any("", err))
			return &pbp.Posts{}, err
		}
		for _, review := range reviews.Reviews {
			post.Reviews = append(post.Reviews, &pbp.Review{
				Id:          review.Id,
				Name:        review.Name,
				Description: review.Description,
				CustomerId:  review.CustomerId,
				Review:      review.Review,
			})
		}

	}
	return posts, nil
}

func (p *PostService) DeletePostByCustomerId(ctx context.Context, req *pbp.Id) (*pbp.IsDeleted, error) {
	exists, err := p.client.Customer().CheckIfCustomerExists(ctx, &pbc.CustomerId{
		Id: req.Id,
	})
	if err != nil {
		p.logger.Error("error -> exists, err := p.client.Customer().CheckIfCustomerExists(ctx, &pbc.CustomerId{", l.Any("error checking customer by id post/service/grpc_client/customer.go", err))
		return &pbp.IsDeleted{}, err
	}
	if !exists.Exists {
		p.logger.Info("There is no such customer")
		return &pbp.IsDeleted{}, status.Error(codes.NotFound, "There is no such customer")
	}

	deleted, ids, err := p.storage.Post().DeletePostByCustomerId(req.Id)
	if err != nil {
		p.logger.Error("error while deleting posts with customer_id grpc_client/post.go", l.Any("", err))
		return &pbp.IsDeleted{
			PostDeleted: false,
		}, err
	}
	if !deleted.PostDeleted {
		return &pbp.IsDeleted{
			PostDeleted: false,
		}, nil
	}
	for _, id := range ids {
		_, err = p.client.Review().DeletePostsReviews(ctx, &pbr.ReviewPostId{
			PostId: id,
		})
		if err != nil {
			p.logger.Error("error while sending id to reviews of post", l.Any("", err))
			return &pbp.IsDeleted{
				PostDeleted: false,
			}, err
		}
	}
	return &pbp.IsDeleted{
		PostDeleted: true,
	}, nil
}

func (p *PostService) GetPostsByPage(ctx context.Context, req *pbp.LimitPage) (*pbp.PostsByPage, error) {
	posts, err := p.storage.Post().GetPostsByPage(req.Page, req.Limit)
	if err != nil {
		p.logger.Error("error while getting posts by page and limit", l.Error(err))
		return &pbp.PostsByPage{}, err
	}
	return posts, nil
}

func (p *PostService) GetPostInfoOnly(ctx context.Context, req *pbp.Id) (*pbp.PostInfoOnly, error){
	postInfo, err := p.storage.Post().GetPostInfoOnly(req.Id)
	if err != nil {
		fmt.Println("Error while sending id to get only post", err)
		return &pbp.PostInfoOnly{}, err
	}
	return postInfo, nil
}
