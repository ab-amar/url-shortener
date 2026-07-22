package service

import (
	"github.com/ab-amar/url-shortener/internal/model"
	"time"
)

func Shorten(originalURL string) model.URL {
	shortenedURL := model.URL{
		OriginalURL: originalURL,
		ShortCode: "short code",
		CreatedAt: time.Now(),
	}
	return shortenedURL
}
