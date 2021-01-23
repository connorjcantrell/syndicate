package postgres

import (
	"fmt"

	"github.com/connorjcantrell/syndicate"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewSeriesStore(db *sqlx.DB) *SeriesStore {
	return &SeriesStore{
		DB: db,
	}
}

type SeriesStore struct {
	*sqlx.DB
}

func (s *SeriesStore) Series(id uuid.UUID) (syndicate.Series, error) {
	var i syndicate.Series
	if err := s.Get(&i, `SELECT * FROM series WHERE id = $1`, id); err != nil {
		return syndicate.Series{}, fmt.Errorf("error getting series: %w", err)
	}
	return i, nil
}

func (s *SeriesStore) Seriess() ([]syndicate.Series, error) {
	var ii []syndicate.Series
	if err := s.Select(&ii, `SELECT * FROM series`); err != nil {
		return []syndicate.Series{}, fmt.Errorf("error getting series: %w", err)
	}
	return ii, nil
}

func (s *SeriesStore) CreateSeries(i *syndicate.Series) error {
	if err := s.Get(i, `INSERT INTO series VALUES ($1, $2, $3, $4) RETURNING *`,
		i.ID,
		i.ManufacturerID,
		i.Name,
		i.SeriesNumber,
		i.Description); err != nil {
		return fmt.Errorf("error creating series: %w", err)
	}
	return nil
}

func (s *SeriesStore) UpdateSeries(i *syndicate.Series) error {
	if err := s.Get(i, `UPDATE series SET name = $1, active = $2, address_id = $3 WHERE id = $4 RETURNING *`,
		i.Name,
		i.ManufacturerID,
		i.SeriesNumber,
		i.Description,
		i.ID); err != nil {
		return fmt.Errorf("error updating series: %w", err)
	}
	return nil
}

func (s *SeriesStore) DeleteSeries(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM series WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting series: %w", err)
	}
	return nil
}
