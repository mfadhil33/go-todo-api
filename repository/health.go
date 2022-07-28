package repository

import "database/sql"

type healthRepository struct {
	db *sql.DB
}

func NewHealthRepository(db *sql.DB) *healthRepository {
	return &healthRepository{db: db}
}

func (r *healthRepository) Ping() error {
	if err := r.db.Ping(); err != nil {
		return err
	}

	return nil
}
