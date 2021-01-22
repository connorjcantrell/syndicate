package postgres

import (
	"fmt"

	"github.com/connorjcantrell/syndicate"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewManufacturerStore(db *sqlx.DB) *ManufacturerStore {
	return &ManufacturerStore{
		DB: db,
	}
}

type ManufacturerStore struct {
	*sqlx.DB
}

func (s *ManufacturerStore) Manufacturer(id uuid.UUID) (syndicate.Manufacturer, error) {
	var m syndicate.Manufacturer
	if err := s.Get(&m, `SELECT * FROM manufacturers WHERE id = $1`, id); err != nil {
		return syndicate.Manufacturer{}, fmt.Errorf("error getting thread: %w", err)
	}
	return m, nil
}

func (s *ManufacturerStore) Manufacturers() ([]syndicate.Manufacturer, error) {
	var mm []syndicate.Manufacturer
	if err := s.Select(&mm, `SELECT * FROM manufacturers`); err != nil {
		return []syndicate.Manufacturer{}, fmt.Errorf("error getting manufacturers: %w", err)
	}
	return mm, nil
}

func (s *ManufacturerStore) CreateManufacturer(m *syndicate.Manufacturer) error {
	if err := s.Get(m, `INSERT INTO manufacturers VALUES ($1, $2, $3, $4) RETURNING *`,
		m.ID,
		m.Name,
		m.Active,
		m.AddressID); err != nil {
		return fmt.Errorf("error creating thread: %w", err)
	}
	return nil
}

func (s *ManufacturerStore) UpdateManufacturer(m *syndicate.Manufacturer) error {
	if err := s.Get(m, `UPDATE manufacturers SET name = $1, active = $2, address_id = $3 WHERE id = $4 RETURNING *`,
		m.Name,
		m.Active,
		m.AddressID,
		m.ID); err != nil {
		return fmt.Errorf("error updating manufacturer: %w", err)
	}
	return nil
}

func (s *ManufacturerStore) DeleteManufacturer(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM manufacturers WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting thread: %w", err)
	}
	return nil
}
