package repository

import (
	"github.com/ab-amar/url-shortener/internal/model"
)

type URLRepository interface {
	SaveURL(url model.URL)
	FindByCode(code string) (model.URL, bool)
}

type InMemoryRepository struct {
	Urls []model.URL
}

func (r *InMemoryRepository) SaveURL(url model.URL) {
	r.Urls = append(r.Urls, url)
}

func (r *InMemoryRepository) FindByCode(code string) (model.URL, bool) {
	urls := r.Urls
	for _, value := range urls {
		if value.ShortCode == code {
			return value, true
		}
	}
	return model.URL{}, false
}
