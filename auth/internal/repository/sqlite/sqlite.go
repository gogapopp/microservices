package sqlite

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	db *sql.DB
}

func NewRepostiry(ctx context.Context, DSN string) (*Repository, error) {
	db, err := sql.Open("sqlite3", DSN)
	if err != nil {
		return &Repository{}, err
	}
	err = db.PingContext(ctx)
	if err != nil {
		return &Repository{}, err
	}
	_, err = db.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username text not null,
		email text not null,
		roles text not null,
		created_at timestamptz,
		updated_at timestamptz not null
	);
	`)
	if err != nil {
		return &Repository{}, err
	}
	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}
