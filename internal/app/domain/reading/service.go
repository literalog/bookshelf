package reading

import (
	"context"
	"time"

	"github.com/literalog/bookshelf/pkg/models"
)

type Service interface {
	Create(ctx context.Context, r *models.Reading) error
	Update(ctx context.Context, r *models.Reading) error
	Delete(ctx context.Context, id string) error
	GetByUserId(ctx context.Context, userId string) ([]models.Reading, error)
	GetByBookId(ctx context.Context, bookId string) ([]models.Reading, error)
	GetById(ctx context.Context, id string) (*models.Reading, error)
	GetAll(ctx context.Context) ([]models.Reading, error)
}

type service struct {
	repository Repository
	validator  Validator
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) Create(ctx context.Context, r *models.Reading) error {
	if r.CreatedAt.IsZero() {
		r.CreatedAt = time.Now()
	}

	if err := s.validator.Validate(r); err != nil {
		return err
	}

	return s.repository.Create(ctx, r)
}

func (s *service) Update(ctx context.Context, r *models.Reading) error {
	return s.repository.Update(ctx, r)
}

func (s *service) Delete(ctx context.Context, id string) error {
	if id == "" {
		return ErrEmptyId
	}
	return s.repository.Delete(ctx, id)
}

func (s *service) GetByUserId(ctx context.Context, userId string) ([]models.Reading, error) {
	if userId == "" {
		return nil, ErrEmptyId
	}
	return s.repository.GetByUserId(ctx, userId)
}

func (s *service) GetByBookId(ctx context.Context, bookId string) ([]models.Reading, error) {
	if bookId == "" {
		return nil, ErrEmptyId
	}
	return s.repository.GetByBookId(ctx, bookId)
}

func (s *service) GetById(ctx context.Context, id string) (*models.Reading, error) {
	if id == "" {
		return nil, ErrEmptyId
	}
	return s.repository.GetById(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]models.Reading, error) {
	return s.repository.GetAll(ctx)
}
