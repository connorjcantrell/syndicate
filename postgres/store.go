package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Store{
		ManufacturerStore: &ManufacturerStore{DB: db},
		SeriesStore:       &SeriesStore{DB: db},
		ModelStore:        &ModelStore{DB: db},
		DistributorStore:  &DistributorStore{DB: db},
		BranchStore:       &BranchStore{DB: db},
	}, nil
}

type Store struct {
	*ManufacturerStore
	*SeriesStore
	*ModelStore
	*DistributorStore
	*BranchStore
}
