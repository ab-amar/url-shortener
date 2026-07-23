package service

import (
	"github.com/ab-amar/url-shortener/internal/model"
	"github.com/ab-amar/url-shortener/internal/repository"
	"time"
	"crypto/sha256"
	"encoding/hex"
)

type URLService interface {
	Shorten(originalURL string) model.URL
}

type ShortenerService struct{
	URLRepo repository.URLRepository
}

func (s ShortenerService) Shorten(originalURL string) model.URL {
	hash := sha256.Sum256([]byte(originalURL))
	hexString := hex.EncodeToString(hash[:])
	shortenedURL := model.URL{
		OriginalURL: originalURL,
		ShortCode: hexString[:6],
		CreatedAt: time.Now(),
	}
	s.URLRepo.SaveURL(shortenedURL)
	return shortenedURL
}
