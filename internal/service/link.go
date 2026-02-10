package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/m-bromo/atom-ly/internal/domain/entities"
	"github.com/m-bromo/atom-ly/internal/hasher"
	repository "github.com/m-bromo/atom-ly/internal/repository/link"
)

type LinkService interface {
	ShortenLink(ctx context.Context, url string) (string, error)
	Redirect(ctx context.Context, shortLink string) (string, error)
}

type linkService struct {
	linkRepository repository.LinkRepository
	hasher         hasher.Hasher
}

func NewLinkService(
	linkRepository repository.LinkRepository,
	hahser hasher.Hasher,
) LinkService {
	return &linkService{
		linkRepository: linkRepository,
		hasher:         hahser,
	}
}

func (s *linkService) ShortenLink(ctx context.Context, url string) (string, error) {
	foundID, err := s.linkRepository.GetByUrl(ctx, url)
	if err != nil && !errors.Is(err, repository.ErrLinkNotFound) {
		return "", fmt.Errorf("failed to shorten link: %w", err)
	}

	if !errors.Is(err, repository.ErrLinkNotFound) {
		code, err := s.hasher.Encode(foundID)
		if err != nil {
			return "", fmt.Errorf("failed to shorten link: %w", err)
		}

		return code, nil
	}

	id, err := s.linkRepository.Save(ctx, &entities.Link{
		Url:       url,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return "", fmt.Errorf("failed to shorten link: %w", err)
	}

	shortCode, err := s.hasher.Encode(id)
	if err != nil {
		return "", fmt.Errorf("failed to shorten link: %w", err)
	}

	return shortCode, nil
}

func (s *linkService) Redirect(ctx context.Context, shortCode string) (string, error) {
	id, err := s.hasher.Decode(shortCode)
	if err != nil {
		return "", fmt.Errorf("failed to redirect link: %w", err)
	}

	url, err := s.linkRepository.GetByID(ctx, id)
	if err != nil && !errors.Is(err, repository.ErrLinkNotFound) {
		return "", fmt.Errorf("failed to redirect link: %w", err)
	}

	if errors.Is(err, repository.ErrLinkNotFound) {
		return url, fmt.Errorf("failed to redirect link: %w", err)
	}

	return url, nil
}
