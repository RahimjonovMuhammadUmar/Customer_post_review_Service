package repo

import (
	pbr "exam/review_service/genproto/review"
)

// ReviewStorage
type ReviewStorage interface {
	CreateReview(*pbr.ReviewRequest) (*pbr.Review, error)
	GetReviews(*pbr.ReviewPostId) (*pbr.Reviews, error)
	DeletePostsReviews(*pbr.ReviewPostId) (*pbr.Empty, error)
	// PostReviews(*pbr.ReviewPostId) (*pbr.Reviews, error)
	DeleteReview(*pbr.ReviewId) (*pbr.Empty, error)

}