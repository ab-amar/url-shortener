package service

import (
	"github.com/ab-amar/url-shortener/internal/model"
	"github.com/ab-amar/url-shortener/internal/repository"
	"time"
)

type URLService interface {
	Shorten(originalURL string) model.URL
}

type ShortenerService struct{
	URLRepo repository.URLRepository
}

func (s ShortenerService) Shorten(originalURL string) model.URL {
	shortenedURL := model.URL{
		OriginalURL: originalURL,
		ShortCode: "short code",
		CreatedAt: time.Now(),
	}
	s.URLRepo.SaveURL(shortenedURL)
	return shortenedURL
}
