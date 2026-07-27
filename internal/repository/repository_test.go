package repository

import (
	"testing"

	"github.com/ab-amar/url-shortener/internal/model"
)

func TestInMemoryRepository_SaveURL(t *testing.T) {
	repo := &InMemoryRepository{}

	url := model.URL{
		OriginalURL: "https://example.com",
		ShortCode:   "abcd1234",
	}

	repo.SaveURL(url)

	if len(repo.Urls) != 1 {
		t.Fatalf("Wrong number of urls")
	}

	if repo.Urls[0] != url {
		t.Fatalf("Wrong url")
	}
}

func TestInMemoryRepository_FindByCodeFound(t *testing.T) {
	repo := &InMemoryRepository{}

	url := model.URL{
		OriginalURL: "https://example.com",
		ShortCode:   "abcd1234",
	}

	repo.SaveURL(url)

	result, found := repo.FindByCode("abcd1234")

	if !found {
		t.Fatalf("Wrong found")
	}
	if result.OriginalURL != url.OriginalURL || result.ShortCode != url.ShortCode {
		t.Fatalf("Wrong url")
	}
}

func TestInMemoryRepository_FindByCodeNotFound(t *testing.T) {
	repo := InMemoryRepository{}

	url := model.URL{
		OriginalURL: "https://example.com",
		ShortCode:   "abcd1234",
	}

	repo.SaveURL(url)

	_, found := repo.FindByCode("xyz")

	if found {
		t.Fatalf("Wrong found")
	}
}
