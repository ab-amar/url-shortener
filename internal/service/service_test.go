package service

import (
	"testing"

	"github.com/ab-amar/url-shortener/internal/model"
	"github.com/stretchr/testify/assert"
)

type fakeRepository struct {
	saveCalled       bool
	savedURL         model.URL
	findByCodeCalled bool
	requestedCode    string
	findResult       model.URL
	findFound        bool
}

func (f *fakeRepository) SaveURL(url model.URL) error {
	f.saveCalled = true
	f.savedURL = url
	return nil
}

func (f *fakeRepository) FindByCode(code string) (model.URL, bool) {
	f.findByCodeCalled = true
	f.requestedCode = code
	return f.findResult, f.findFound
}

func TestShortenerService_ResolveFound(t *testing.T) {
	code := "abcd1234"
	f := fakeRepository{
		findResult: model.URL{
			OriginalURL: "https://example.com",
			ShortCode:   code,
		},
		findFound: true,
	}
	s := ShortenerService{
		URLRepo: &f,
	}

	result, found := s.Resolve(code)

	assert.True(t, f.findByCodeCalled)
	assert.Equal(t, code, f.requestedCode)
	assert.Equal(t, f.findResult.OriginalURL, result.OriginalURL)
	assert.Equal(t, f.findResult.ShortCode, result.ShortCode)
	assert.Equal(t, f.findFound, found)
}

func TestShortenerService_ResolveNotFound(t *testing.T) {
	code := "abcd1234"
	f := fakeRepository{
		findResult: model.URL{},
		findFound:  false,
	}
	s := ShortenerService{
		URLRepo: &f,
	}

	_, found := s.Resolve(code)

	assert.True(t, f.findByCodeCalled)
	assert.False(t, found)
	assert.Equal(t, code, f.requestedCode)
}
func TestShortenerService_Shorten(t *testing.T) {
	originalURL := "https://example.com"
	f := fakeRepository{}
	s := ShortenerService{
		URLRepo: &f,
	}
	result, _ := s.Shorten(originalURL)

	assert.Equal(t, originalURL, result.OriginalURL)
	assert.NotEmpty(t, result.ShortCode)
	assert.NotZero(t, result.CreatedAt)
	assert.True(t, f.saveCalled)
	assert.Equal(t, result.OriginalURL, f.savedURL.OriginalURL)
	assert.Equal(t, result.ShortCode, f.savedURL.ShortCode)
	assert.Equal(t, result.CreatedAt, f.savedURL.CreatedAt)
}
