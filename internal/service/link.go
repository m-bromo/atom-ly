package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/m-bromo/atom-ly/internal/domain/entities"
	"github.com/m-bromo/atom-ly/internal/domain/repository"
	"github.com/m-bromo/atom-ly/internal/hasher"
)

type LinkService interface {
	ShortenLink(ctx context.Context, url string) (string, error)
	Redirect(ctx context.Context, shortLink string) (string, error)
}

type linkService struct {
	linkRepository repository.LinkRepository
	hasher         hasher.Hahser
}

func NewLinkService(
	linkRepository repository.LinkRepository,
	hahser hasher.Hahser,
) LinkService {
	return &linkService{
		linkRepository: linkRepository,
		hasher:         hahser,
	}
}

func (s *linkService) ShortenLink(ctx context.Context, url string) (string, error) {
	link, err := s.linkRepository.GetByUrl(ctx, url)
	if err != nil {
		slog.Error("failed to get link", "error", err.Error())
		return "", err
	}

	if link != nil {
		return link.ShortCode, nil
	}

	id, err := s.linkRepository.Save(ctx, &entities.Link{
		Url:       url,
		CreatedAt: time.Now(),
	})
	if err != nil {
		slog.Error("failed to save link", "error", err.Error())
		return "", err
	}

	shortCode, err := s.hasher.Encode(id)
	if err != nil {
		slog.Error("failed to hash id", "error", err.Error())
		return "", err
	}

	return shortCode, nil
}

func (s *linkService) Redirect(ctx context.Context, shortCode string) (string, error) {
	id, err := s.hasher.Decode(shortCode)
	if err != nil {
		slog.Error("failed to decode short code", "error", err.Error())
		return "", err
	}

	url, err := s.linkRepository.GetByID(ctx, id)
	if err != nil {
		slog.Error("failed to get url", "error", err.Error())
		return "", err
	}

	return url, nil
}
