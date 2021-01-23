package postgres

import (
	"fmt"

	"github.com/connorjcantrell/syndicate"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type DistributorStore struct {
	*sqlx.DB
}

func (s *DistributorStore) Distributor(id uuid.UUID) (syndicate.Distributor, error) {
	var d syndicate.Distributor
	if err := s.Get(&d,
		`SELECT * FROM distributors 
		WHERE id = $1`, id); err != nil {
		return syndicate.Distributor{}, fmt.Errorf("error getting distributor: %w", err)
	}
	return d, nil
}

func (s *DistributorStore) Distributors() ([]syndicate.Distributor, error) {
	var dd []syndicate.Distributor
	if err := s.Get(&dd,
		`SELECT * FROM distributors
		ORDER BY name`); err != nil {
		return []syndicate.Distributor{}, fmt.Errorf("error gettting distributors: %w", err)
	}
	return dd, nil
}

func (s *DistributorStore) CreateDistributor(d *syndicate.Distributor) error {
	if err := s.Get(d,
		`INSERT INTO distributors
		 VALUES ($1, $2, $3, $4) 
		 RETURNING *`,
		d.ID,
		d.Active,
		d.Name,
		d.AddressID); err != nil {
		return fmt.Errorf("error creating distributor: %w", err)
	}
	return nil
}

func (s *DistributorStore) UpdateDistributor(d *syndicate.Distributor) error {
	if err := s.Get(d,
		`UPDATE distributors 
		SET name = $1, active = $2, address_id = $3 
		WHERE id = $4 RETURNING *`,
		d.Name,
		d.Active,
		d.AddressID,
		d.ID); err != nil {
		return fmt.Errorf("error updating distributors: %w", err)
	}
	return nil
}

func (s *DistributorStore) DeleteDistributor(id uuid.UUID) error {
	if _, err := s.Exec(`
	DELETE FROM distributors 
	WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting distributor: %w", err)
	}
	return nil
}
