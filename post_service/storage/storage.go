package storage

import (
	"exam/post_service/storage/repo"
	"exam/post_service/storage/postgres"

	"github.com/jmoiron/sqlx"
)

// IStorage ...
type IStorage interface {
	Post() repo.PostStorage
}

type storagePg struct {
	db          *sqlx.DB
	postRepo repo.PostStorage
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:          db,
		postRepo: postgres.NewProductRepo(db),
	}
}

func (s storagePg) Post() repo.PostStorage {
	return s.postRepo
}
