package storage

import (
	"exam/customer_service/storage/postgres"
	"exam/customer_service/storage/repo"
	

	"github.com/jmoiron/sqlx"
)

// IStorage ...
type IStorage interface {
	Customer() repo.CustomerStorage
}

type storagePg struct {
	db          *sqlx.DB
	customerRepo repo.CustomerStorage
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:          db,
		customerRepo: postgres.NewCustomerRepo(db),
	}
}

func (s storagePg) Customer() repo.CustomerStorage {
	return s.customerRepo
}
