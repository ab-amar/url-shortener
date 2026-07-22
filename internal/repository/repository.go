package repository

import (
	"github.com/ab-amar/url-shortener/internal/model"
)

type URLRepository interface {
	SaveURL(url model.URL)
}

type InMemoryRepository struct {
	Urls []model.URL
}

func (r *InMemoryRepository) SaveURL(url model.URL) {
	r.Urls = append(r.Urls, url)
}
