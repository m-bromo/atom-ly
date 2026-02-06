package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/m-bromo/atom-ly/internal/database/postgres/sqlc"
	"github.com/m-bromo/atom-ly/internal/domain/entities"
)

type postgresLinkRepository struct {
	querier sqlc.Querier
}

func NewPostgresLinkRepository(querier sqlc.Querier) LinkRepository {
	return &postgresLinkRepository{
		querier: querier,
	}
}

func (r *postgresLinkRepository) Save(ctx context.Context, link *entities.Link) (int, error) {
	id, err := r.querier.Save(ctx, sqlc.SaveParams{
		Url: link.Url,
		CreatedAt: sql.NullTime{
			Time:  link.CreatedAt,
			Valid: true,
		},
	})

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *postgresLinkRepository) GetByID(ctx context.Context, id int) (string, error) {
	url, err := r.querier.GetByID(ctx, int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}

		return "", err
	}

	return url, nil
}
func (r *postgresLinkRepository) GetByUrl(ctx context.Context, url string) (int, error) {
	id, err := r.querier.GetIDByUrl(ctx, url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}

		return 0, err
	}

	return int(id), nil
}
