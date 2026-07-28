package service

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/ab-amar/url-shortener/internal/model"
	"github.com/ab-amar/url-shortener/internal/repository"
)

type URLService interface {
	Shorten(originalURL string) (model.URL, error)
	Resolve(code string) (model.URL, bool)
}

type ShortenerService struct {
	URLRepo repository.URLRepository
}

func (s ShortenerService) Shorten(originalURL string) (model.URL, error) {
	hash := sha256.Sum256([]byte(originalURL))
	hexString := hex.EncodeToString(hash[:])
	shortenedURL := model.URL{
		OriginalURL: originalURL,
		ShortCode:   hexString[:8],
		CreatedAt:   time.Now(),
	}
	err := s.URLRepo.SaveURL(shortenedURL)
	if err != nil {
		return model.URL{}, err
	}
	return shortenedURL, err
}

func (s ShortenerService) Resolve(code string) (model.URL, bool) {
	return s.URLRepo.FindByCode(code)
}
