package service

import (
	"context"
	"errors"
	"time"

	"github.com/rseigha/goecomapi/internal/domain"
	"github.com/rseigha/goecomapi/internal/repository"
)

type OrderService interface {
	CreateOrder(ctx context.Context, o *domain.Order) error
	GetByID(ctx context.Context, id string) (*domain.Order, error)
	GetByUser(ctx context.Context, userID string) ([]*domain.Order, error)
}

type orderService struct {
	repo repository.OrderRepository
	// optionally productRepo to validate stock etc.
}

func NewOrderService(r repository.OrderRepository) OrderService {
	return &orderService{repo: r}
}

func (s *orderService) CreateOrder(ctx context.Context, o *domain.Order) error {
	if len(o.Items) == 0 {
		return errors.New("order must contain items")
	}
	total := 0.0
	for _, it := range o.Items {
		total += it.Price * float64(it.Quantity)
	}
	o.Total = total
	o.Status = domain.OrderPending
	o.CreatedAt = time.Now().UTC()
	o.UpdatedAt = o.CreatedAt
	return s.repo.Create(ctx, o)
}

func (s *orderService) GetByID(ctx context.Context, id string) (*domain.Order, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *orderService) GetByUser(ctx context.Context, userID string) ([]*domain.Order, error) {
	return s.repo.GetByUserID(ctx, userID)
}
