package service

import (
	"testing"

	"github.com/ab-amar/url-shortener/internal/model"
)

type fakeRepository struct {
	saveCalled       bool
	savedURL         model.URL
	findByCodeCalled bool
	requestedCode    string
	findResult       model.URL
	findFound        bool
}

func (f *fakeRepository) SaveURL(url model.URL) {
	f.saveCalled = true
	f.savedURL = url
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

	if !f.findByCodeCalled {
		t.Fatalf(" findByCodeCalled not called")
	}
	if result.OriginalURL != f.findResult.OriginalURL {
		t.Fatalf("wrong original url")
	}

	if result.ShortCode != f.findResult.ShortCode {
		t.Fatalf("wrong short code")
	}
	if f.requestedCode != code {
		t.Fatalf("Wrong code")
	}

	if f.findFound != found {
		t.Fatalf("Wrong found")
	}
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

	if !f.findByCodeCalled {
		t.Fatalf(" findByCodeCalled not called")
	}
	if found {
		t.Fatalf("Wrong found")
	}

	if f.requestedCode != code {
		t.Fatalf("Wrong code")
	}
}
func TestShortenerService_Shorten(t *testing.T) {
	originalURL := "https://example.com"
	f := fakeRepository{}
	s := ShortenerService{
		URLRepo: &f,
	}
	result := s.Shorten(originalURL)

	if result.OriginalURL != originalURL {
		t.Fatalf("Wrong original url")
	}

	if result.ShortCode == "" {
		t.Fatalf("Empty short code")
	}

	if result.CreatedAt.IsZero() {
		t.Fatalf("Empty timestamp")
	}

	if !f.saveCalled {
		t.Fatalf("Save not called")
	}
	if f.savedURL.OriginalURL != result.OriginalURL {
		t.Fatalf("Saved url wrong")
	}
	if f.savedURL.ShortCode != result.ShortCode {
		t.Fatalf("Saved code wrong")
	}
	if f.savedURL.CreatedAt != result.CreatedAt {
		t.Fatalf("Created At wrong")
	}
}
