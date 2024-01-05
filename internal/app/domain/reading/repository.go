package reading

import (
	"context"

	"github.com/literalog/bookshelf/pkg/models"
)

type Repository interface {
	Create(ctx context.Context, r *models.Reading) error
	Update(ctx context.Context, r *models.Reading) error
	Delete(ctx context.Context, id string) error
	GetByUserId(ctx context.Context, id string) ([]models.Reading, error)
	GetByBookId(ctx context.Context, id string) ([]models.Reading, error)
	GetById(ctx context.Context, id string) (*models.Reading, error)
	GetAll(ctx context.Context) ([]models.Reading, error)
}
