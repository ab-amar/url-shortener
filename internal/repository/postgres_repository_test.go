package repository

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ab-amar/url-shortener/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestPostgresRepository(t *testing.T) *PostgresRepository {
	t.Helper()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Skip("DATABASE_URL is not set")
	}

	pool, err := pgxpool.New(context.Background(), databaseURL)
	require.NoError(t, err)

	t.Cleanup(func() {
		pool.Close()
	})

	_, err = pool.Exec(context.Background(), "TRUNCATE TABLE urls RESTART IDENTITY;")
	require.NoError(t, err)

	t.Cleanup(func() {
		_, err := pool.Exec(context.Background(), "TRUNCATE TABLE urls RESTART IDENTITY;")
		require.NoError(t, err)
	})

	repo := NewPostgresRepository(pool)
	return &repo
}

func TestPostgresRepository_SaveURLAndFindByCode(t *testing.T) {
	repo := newTestPostgresRepository(t)
	now := time.Now().UTC().Truncate(time.Microsecond)
	expiresAt := now.Add(24 * time.Hour).Truncate(time.Microsecond)

	url := model.URL{
		OriginalURL: "https://example.com",
		ShortCode:   "pgtest01",
		CreatedAt:   now,
		UpdatedAt:   now,
		ExpiresAt:   &expiresAt,
	}

	err := repo.SaveURL(url)
	require.NoError(t, err)

	result, found := repo.FindByCode("pgtest01")

	assert.True(t, found)
	assert.Equal(t, url.OriginalURL, result.OriginalURL)
	assert.Equal(t, url.ShortCode, result.ShortCode)
	assert.WithinDuration(t, url.CreatedAt, result.CreatedAt, time.Second)
	assert.WithinDuration(t, url.UpdatedAt, result.UpdatedAt, time.Second)
	require.NotNil(t, result.ExpiresAt)
	assert.WithinDuration(t, expiresAt, *result.ExpiresAt, time.Second)
}

func TestPostgresRepository_FindByCodeNotFound(t *testing.T) {
	repo := newTestPostgresRepository(t)

	result, found := repo.FindByCode("missing01")

	assert.False(t, found)
	assert.Equal(t, model.URL{}, result)
}
