package model

import (
	"time"
)

type URL struct {
	OriginalURL string     `json:"original_url"`
	ShortCode   string     `json:"short_code"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	ExpiresAt   *time.Time `json:"expires_at"`
}
