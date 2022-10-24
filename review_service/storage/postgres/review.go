package postgres

import (
	pbr "exam/review_service/genproto/review"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type reviewRepo struct {
	db *sqlx.DB
}

// NewReviewRepo ...
func NewReviewRepo(db *sqlx.DB) *reviewRepo {
	return &reviewRepo{
		db: db,
	}
}

func (r reviewRepo) CreateReview(req *pbr.ReviewRequest) (*pbr.Review, error) {
	rev := &pbr.Review{}
	err := r.db.QueryRow(`INSERT INTO ratings(
		name, 
		description, 
		rating, 
		customer_id, 
		post_id) VALUES($1, $2, $3, $4, $5) RETURNING id, name, description, rating, customer_id, post_id`,
		req.Name,
		req.Description,
		req.Review,
		req.CustomerId,
		req.PostId).Scan(
		&rev.Id,
		&rev.Name,
		&rev.Description,
		&rev.Review,
		&rev.CustomerId,
		&rev.PostId,
	)
	if err != nil {
		fmt.Println("error while inserting into ratings")
		return &pbr.Review{}, err
	}

	return rev, nil
}
func (r reviewRepo) GetReviews(req *pbr.ReviewPostId) (*pbr.Reviews, error) {
	reviews_list := &pbr.Reviews{}
	reviews, err := r.db.Query(`SELECT 
	id, 
	name, 
	description, 
	rating, 
	customer_id FROM ratings WHERE post_id = $1`, req.PostId)
	if err != nil {
		fmt.Println("error while selecting reviews from ratings", err)
		return &pbr.Reviews{}, err
	}

	for reviews.Next() {
		review := &pbr.Review{}
		err = reviews.Scan(
			&review.Id,
			&review.Name,
			&review.Description,
			&review.Review,
			&review.CustomerId,
		)
		if err != nil {
			fmt.Println("error while scanning to review from reviews", err)
			return &pbr.Reviews{}, err
		}
		reviews_list.Reviews = append(reviews_list.Reviews, review)
	}

	return reviews_list, nil
}
func (r reviewRepo) DeleteReview(req *pbr.ReviewPostId) (*pbr.Empty, error) {
	_, err := r.db.Exec(`UPDATE ratings SET deleted_at = NOW() WHERE post_id = $1`, req.PostId)
	if err != nil {
		fmt.Println("Error while deleting from ratings", err)
		return &pbr.Empty{}, err
	}

	return &pbr.Empty{}, nil
}

func (r reviewRepo) PostReviews(req *pbr.ReviewPostId) (*pbr.Reviews, error) {
	reviews := &pbr.Reviews{}
	reviews_rows, err := r.db.Query(`SELECT id, name, description, customer_id, rating FROM ratings WHERE post_id = $1`, req.PostId)
	if err != nil {
		fmt.Println("Error while selecting from ratings", err)
		return &pbr.Reviews{}, err
	}
	
	for reviews_rows.Next() {
		review := &pbr.Review{}
		err = reviews_rows.Scan(
			&review.Id,
			&review.Name,
			&review.Description,
			&review.CustomerId,
			&review.Review,
		)
		if err != nil {
			fmt.Println("error scanning to review", err)
			return &pbr.Reviews{}, err
		}
		reviews.Reviews = append(reviews.Reviews, review)
	}
	return reviews, nil
}
