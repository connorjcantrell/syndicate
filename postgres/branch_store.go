package postgres

import (
	"fmt"

	"github.com/connorjcantrell/syndicate"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type BranchStore struct {
	*sqlx.DB
}

func (s *BranchStore) Branch(id uuid.UUID) (syndicate.Branch, error) {
	var b syndicate.Branch
	if err := s.Get(&b,
		`SELECT * FROM branches
		WHERE id = $1`, id); err != nil {
		return syndicate.Branch{}, fmt.Errorf("error getting branch: %w", err)
	}
	return b, nil
}

func (s *BranchStore) Branches() ([]syndicate.Branch, error) {
	var bb []syndicate.Branch
	if err := s.Get(&bb,
		`SELECT * FROM branches
		ORDER BY name`); err != nil {
		return []syndicate.Branch{}, fmt.Errorf("error getting branches: %w", err)
	}
	return bb, nil
}

func (s *BranchStore) BranchesByDistributor(id uuid.UUID) ([]syndicate.Branch, error) {
	var bb []syndicate.Branch
	if err := s.Get(&bb,
		`SELECT * FROM branches
		WHERE id = $1
		ORDER BY name`, id); err != nil {
		return []syndicate.Branch{}, fmt.Errorf("error getting branches: %w", err)
	}
	return bb, nil
}
func (s *BranchStore) CreateBranch(b *syndicate.Branch) error {
	if err := s.Get(b, `INSERT INTO branches
	VALUES ($1, $2, $3, $4)
	RETURNING *`,
		b.ID,
		b.DistributorID,
		b.AddressID,
		b.Active); err != nil {
		return fmt.Errorf("error creating branch: %w", err)
	}
	return nil
}

func (s *BranchStore) UpdateBranch(b *syndicate.Branch) error {
	if err := s.Get(b,
		`UPDATE branches
		SET distributor_id = $2, address_id = $3, active = $4
		WHERE id = $1
		RETURNING *`,
		b.ID,
		b.DistributorID,
		b.AddressID,
		b.Active); err != nil {
		return fmt.Errorf("error updating branch: %w", err)
	}
	return nil
}

func (s *BranchStore) DeleteBranch(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM branches WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting branch: %w", err)
	}
	return nil
}
