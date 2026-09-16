package repository

import (
	"context"

	"url-shortener-go/internal/domain"
)

type URLRepository interface {
	Save(ctx context.Context, url domain.URL) error
	FindByShortCode(ctx context.Context, shortCode string) (domain.URL, error)
}
