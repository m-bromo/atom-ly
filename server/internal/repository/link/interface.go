package repository

import (
	"context"
	"errors"

	"github.com/m-bromo/atom-ly/internal/domain/entities"
)

var (
	ErrLinkNotFound = errors.New("link not found in database")
)

type LinkRepository interface {
	Save(ctx context.Context, link *entities.Link) (int, error)
	GetByID(ctx context.Context, id int) (string, error)
	GetByUrl(ctx context.Context, url string) (int, error)
}
