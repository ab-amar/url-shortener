package service

import (
	"github.com/ab-amar/url-shortener/internal/model"
	"time"
)

type URLService interface {
	Shorten(originalURL string) model.URL
}

type ShortenerService struct{}

func (s ShortenerService) Shorten(originalURL string) model.URL {
	shortenedURL := model.URL{
		OriginalURL: originalURL,
		ShortCode: "short code",
		CreatedAt: time.Now(),
	}
	return shortenedURL
}
