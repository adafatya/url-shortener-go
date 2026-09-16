package memory

import (
	"context"
	"errors"
	"sync"

	"url-shortener-go/internal/domain"
)

var ErrURLNotFound = errors.New("url not found")

type URLRepository struct {
	mu   sync.RWMutex
	data map[string]domain.URL
}

func NewURLRepository() *URLRepository {
	return &URLRepository{data: make(map[string]domain.URL)}
}

func (r *URLRepository) Save(_ context.Context, url domain.URL) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[url.ShortCode] = url
	return nil
}

func (r *URLRepository) FindByShortCode(_ context.Context, shortCode string) (domain.URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	url, ok := r.data[shortCode]
	if !ok {
		return domain.URL{}, ErrURLNotFound
	}
	return url, nil
}
