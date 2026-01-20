package service

import (
	"context"

	"github.com/rseigha/goecomapi/internal/domain"
	"github.com/rseigha/goecomapi/internal/repository"
)

type ProductService interface {
	Create(ctx context.Context, p *domain.Product) error
	GetByID(ctx context.Context, id string) (*domain.Product, error)
	Update(ctx context.Context, p *domain.Product) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, page int) ([]*domain.Product, int64, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return &productService{repo: r}
}

func (s *productService) Create(ctx context.Context, p *domain.Product) error {
	return s.repo.Create(ctx, p)
}

func (s *productService) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *productService) Update(ctx context.Context, p *domain.Product) error {
	return s.repo.Update(ctx, p)
}

func (s *productService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *productService) List(ctx context.Context, limit, page int) ([]*domain.Product, int64, error) {
	return s.repo.List(ctx, limit, page)
}
