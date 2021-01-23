package postgres

import (
	"fmt"

	"github.com/connorjcantrell/syndicate"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewModelStore(db *sqlx.DB) *ModelStore {
	return &ModelStore{
		DB: db,
	}
}

type ModelStore struct {
	*sqlx.DB
}

func (s *ModelStore) Model(id uuid.UUID) (syndicate.Model, error) {
	var m syndicate.Model
	if err := s.Get(&m, `SELECT * FROM models WHERE id = $1`, id); err != nil {
		return syndicate.Model{}, fmt.Errorf("error getting model: %w", err)
	}
	return m, nil
}

func (s *ModelStore) ModelsByManufacturer(manufacturerID uuid.UUID) ([]syndicate.Model, error) {
	var mm []syndicate.Model
	if err := s.Get(&mm, `SELECT * FROM models WHERE manufacturer_id = $1`, manufacturerID); err != nil {
		return []syndicate.Model{}, fmt.Errorf("error getting models: %w", err)
	}
	return mm, nil
}

func (s *ModelStore) ModelsBySeries(seriesID uuid.UUID) ([]syndicate.Model, error) {
	var mm []syndicate.Model
	if err := s.Get(&mm, `SELECT * FROM models WHERE series_id = $1`, seriesID); err != nil {
		return []syndicate.Model{}, fmt.Errorf("error getting models: %w", err)
	}
	return mm, nil
}
func (s *ModelStore) CreateModel(m *syndicate.Model) error {
	if err := s.Get(m, `INSERT INTO models VALUES ($1, $2, $3, $4, $5) RETURNING *`,
		m.ID,
		m.ModelNumber,
		m.ManufacturerID,
		m.SeriesID,
		m.Image); err != nil {
		return fmt.Errorf("error creating model: %w", err)
	}
	return nil
}

func (s *ModelStore) UpdateModel(m *syndicate.Model) error {
	if err := s.Get(m, `UPDATE models SET 
		model_number = $1,
		manufacturer_id = $2,
		series_id = $3,
		image = $4,
		WHERE id = $5 RETURNING *`,
		m.ModelNumber,
		m.ManufacturerID,
		m.SeriesID,
		m.Image,
		m.ID); err != nil {
		return fmt.Errorf("error updating model: %w", err)
	}
	return nil
}

func (s *ModelStore) DeleteModel(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM models WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting model: %w", err)
	}
	return nil
}
