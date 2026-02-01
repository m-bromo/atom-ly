package repository

import (
	"context"

	"github.com/m-bromo/atom-ly/internal/domain/entities"
)

type LinkRepository interface {
	Save(ctx context.Context, link *entities.Link) (int, error)
	GetByID(ctx context.Context, id int) (string, error)
	GetByUrl(ctx context.Context, url string) (string, error)
}
