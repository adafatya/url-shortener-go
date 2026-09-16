package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/url"
	"time"

	"url-shortener-go/internal/domain"
	"url-shortener-go/internal/repository"
)

var ErrInvalidTargetURL = errors.New("invalid target url")

type URLService struct {
	repository repository.URLRepository
}

func NewURLService(urlRepository repository.URLRepository) *URLService {
	return &URLService{repository: urlRepository}
}

func (s *URLService) Shorten(ctx context.Context, targetURL string) (domain.URL, error) {
	parsedURL, err := url.ParseRequestURI(targetURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return domain.URL{}, ErrInvalidTargetURL
	}

	shortCode, err := generateShortCode()
	if err != nil {
		return domain.URL{}, err
	}

	shortURL := domain.URL{
		ID:        shortCode,
		ShortCode: shortCode,
		TargetURL: targetURL,
		CreatedAt: time.Now().UTC(),
	}
	if err := s.repository.Save(ctx, shortURL); err != nil {
		return domain.URL{}, err
	}

	return shortURL, nil
}

func (s *URLService) Resolve(ctx context.Context, shortCode string) (domain.URL, error) {
	return s.repository.FindByShortCode(ctx, shortCode)
}

func generateShortCode() (string, error) {
	bytes := make([]byte, 6)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}
