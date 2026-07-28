package repository

import (
	"context"

	"github.com/ab-amar/url-shortener/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	Pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) PostgresRepository {
	return PostgresRepository{
		Pool: pool,
	}
}

func (r *PostgresRepository) SaveURL(url model.URL) error {
	query := "INSERT INTO urls (short_code, original_url, created_at, updated_at, expires_at) VALUES ($1, $2, $3, $4, $5);"
	_, err := r.Pool.Exec(
		context.Background(),
		query,
		url.ShortCode,
		url.OriginalURL,
		url.CreatedAt,
		url.UpdatedAt,
		url.ExpiresAt,
	)
	return err
}

func (r *PostgresRepository) FindByCode(code string) (model.URL, bool) {
	query := "SELECT original_url, short_code, created_at, updated_at, expires_at FROM urls WHERE short_code = $1;"
	row := r.Pool.QueryRow(context.Background(), query, code)
	var url model.URL
	err := row.Scan(
		&url.OriginalURL,
		&url.ShortCode,
		&url.CreatedAt,
		&url.UpdatedAt,
		&url.ExpiresAt,
	)

	if err == pgx.ErrNoRows {
		return model.URL{}, false
	}
	if err != nil {
		return model.URL{}, false
	}
	return url, true
}
