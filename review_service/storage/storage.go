package storage

import (
	"exam/review_service/storage/repo"
	"exam/review_service/storage/postgres"

	"github.com/jmoiron/sqlx"
)

// IStorage ...
type IStorage interface {
	Review() repo.ReviewStorage
}

type storagePg struct {
	db         *sqlx.DB
	reviewRepo repo.ReviewStorage
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db: db,
		reviewRepo: postgres.NewReviewRepo(db),
	}
}
func (s storagePg) Review() repo.ReviewStorage {
	return s.reviewRepo
}