package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/berikarg/fortune-wheel/api"
)

type SpinResultRepository struct {
	db *sqlx.DB
}

func NewSpinResultRepository(db *sqlx.DB) *SpinResultRepository {
	return &SpinResultRepository{db: db}
}

func (r *SpinResultRepository) GetAll() ([]api.SpinResult, error) {
	query := `SELECT * FROM spin_results`
	var spinResults []api.SpinResult
	err := r.db.Select(&spinResults, query)
	if err != nil {
		return nil, err
	}
	return spinResults, nil
}
func (r *SpinResultRepository) Create(result api.SpinResult) error {
	query := `INSERT INTO spin_results (result, time) VALUES ($1, $2)`
	_, err := r.db.Queryx(query, result.Result, result.Timestamp)
	return err
}
