package usecase

import (
	"context"
	"testing"

	"url-shortener-go/internal/repository/memory"
)

func TestURLServiceShortenAndResolve(t *testing.T) {
	service := NewURLService(memory.NewURLRepository())

	created, err := service.Shorten(context.Background(), "https://example.com/path")
	if err != nil {
		t.Fatalf("Shorten() error = %v", err)
	}
	if created.ShortCode == "" {
		t.Fatal("Shorten() returned an empty short code")
	}

	resolved, err := service.Resolve(context.Background(), created.ShortCode)
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	if resolved.TargetURL != created.TargetURL {
		t.Fatalf("Resolve() target = %q, want %q", resolved.TargetURL, created.TargetURL)
	}
}

func TestURLServiceShortenRejectsInvalidURL(t *testing.T) {
	service := NewURLService(memory.NewURLRepository())

	_, err := service.Shorten(context.Background(), "not-a-url")
	if err != ErrInvalidTargetURL {
		t.Fatalf("Shorten() error = %v, want %v", err, ErrInvalidTargetURL)
	}
}
