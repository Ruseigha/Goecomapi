package service

import (
	"context"
	"time"

	"github.com/rseigha/goecomapi/internal/domain"
	"github.com/rseigha/goecomapi/internal/repository"
	"go.uber.org/zap"
)

type UserService interface {
	GetByID(ctx context.Context, id string) (*domain.User, error)
	Update(ctx context.Context, u *domain.User) error
	Delete(ctx context.Context, id string) error
}

type userService struct {
	userRep repository.UserRepository
	logger  *zap.Logger
}

func NewUserService(ur repository.UserRepository, logger *zap.Logger) UserService {
	return &userService{
		userRep: ur,
		logger:  logger,
	}
}

func (s *userService) GetByID(ctx context.Context, id string) (*domain.User, error) {
	u, err := s.userRep.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	u.PasswordHash = ""
	return u, nil
}

func (s *userService) Update(ctx context.Context, u *domain.User) error {
	u.UpdatedAt = time.Now().UTC()
	return s.userRep.Update(ctx, u)
}

func (s *userService) Delete(ctx context.Context, id string) error {
	return s.userRep.Delete(ctx, id)
}
