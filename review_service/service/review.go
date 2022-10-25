package service

import (
	"context"
	"database/sql"
	"fmt"

	pbp "exam/review_service/genproto/post"
	pbr "exam/review_service/genproto/review"
	l "exam/review_service/pkg/logger"
	grpcClient "exam/review_service/service/grpc_client"
	"exam/review_service/storage"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ReviewService ...
type ReviewService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.GrpcClientI
}

// NewReviewService ...
func NewReviewService(db *sqlx.DB, log l.Logger, client grpcClient.GrpcClientI) *ReviewService {
	return &ReviewService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (s *ReviewService) CreateReview(ctx context.Context, req *pbr.ReviewRequest) (*pbr.Review, error) {
	exist, err := s.client.Post().DoesPostExist(ctx, &pbp.Id{
		Id: req.PostId,
	})
	if err != nil {
		if err != nil {
			s.logger.Error("error -> exists, err := s.client.Post().DoesPostExist(ctx, &pbp.Ids{", l.Any("error checking post by id review_service/grpc_client/review.go", err))
			return &pbr.Review{}, err
		}
	}
	if !exist.Exists {
		s.logger.Info("such post does not exist")
		return &pbr.Review{}, nil
	}

	review, err := s.storage.Review().CreateReview(req)
	if err != nil {
		s.logger.Error("error insert", l.Any("error insert review", err))
		return &pbr.Review{}, status.Error(codes.Internal, "something went wrong, please check review")
	}
	return review, nil
}

func (s *ReviewService) GetReviews(ctx context.Context, req *pbr.ReviewPostId) (*pbr.Reviews, error) {
	reviews, err := s.storage.Review().GetReviews(req)
	if err != nil && err != sql.ErrNoRows {
		s.logger.Error("error while sending request to db level GetReviews", l.Any("error while searching for post reviews", err))
		return &pbr.Reviews{}, err
	}
	return reviews, nil
}

func (s *ReviewService) DeletePostsReviews(ctx context.Context, req *pbr.ReviewPostId) (*pbr.Empty, error) {
	_, err := s.storage.Review().DeletePostsReviews(req)
	if err != nil {
		s.logger.Error("error while sending id to delete from ratings", l.Any("", err))
		return &pbr.Empty{}, err
	}
	return &pbr.Empty{}, nil
}

// func (s *ReviewService) PostReviews(ctx context.Context, req *pbr.ReviewPostId) (*pbr.Reviews, error) {
// 	reviews, err := s.storage.Review().PostReviews(req)
// 	if err != nil {
// 		fmt.Println("error while sending to PostReviews", err)
// 		return &pbr.Reviews{}, err
// 	}

// 	return reviews, nil
// }

func (s *ReviewService) DeleteReview(ctx context.Context, req *pbr.ReviewId) (*pbr.Empty, error) {
	_, err := s.storage.Review().DeleteReview(req)
	if err != nil {
		fmt.Println("error while sending to DeletReview", err)
		return &pbr.Empty{}, err
	}
	return &pbr.Empty{}, nil
}

func (s *ReviewService) GetReview(ctx context.Context, req *pbr.ReviewId) (*pbr.Review, error) {
	review, err := s.storage.Review().GetReview(req)
	if err != nil {
		fmt.Println("error service/review.go 98", err)
		return &pbr.Review{}, err
	}

	return review, nil

}
