package service

import (
	"context"
	"errors"
	"time"

	"github.com/rseigha/goecomapi/internal/domain"
	"github.com/rseigha/goecomapi/internal/repository"
	"github.com/rseigha/goecomapi/pkg/hash"
	jwtpkg "github.com/rseigha/goecomapi/pkg/jwt"
	"go.uber.org/zap"
)

type AuthService interface {
	Register(ctx context.Context, name, email, password string) (*domain.User, error)
	Login(ctx context.Context, email, password string) (string, error) // returns JWT token
}

type authService struct {
	userRepo repository.UserRepository
	jwt      *jwtpkg.JWT
	logger   *zap.Logger
}

func NewAuthService(u repository.UserRepository, j *jwtpkg.JWT, logger *zap.Logger) AuthService {
	return &authService{userRepo: u, jwt: j, logger: logger}
}
func (s *authService) Register(ctx context.Context, name, email, password string) (*domain.User, error) {
	// check existing user
	existing, _ := s.userRepo.GetByEmail(ctx, email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}
	hashed, err := hash.HashPassword(password)
	if err != nil {
		return nil, err
	}
	u := &domain.User{
		Name:         name,
		Email:        email,
		PasswordHash: hashed,
		Role:         domain.RoleUser,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}
	if err := s.userRepo.Create(ctx, u); err != nil {
		return nil, err
	}
	u.PasswordHash = "" // redact
	return u, nil
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	u, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	if u == nil {
		return "", errors.New("invalid credentials")
	}
	if err := hash.CheckPassword(u.PasswordHash, password); err != nil {
		return "", errors.New("invalid credentials")
	}
	// generate token
	claims := jwtpkg.CustomClaims{
		UserID: u.ID,
		Email:  u.Email,
		Role:   string(u.Role),
	}
	token, err := s.jwt.GenerateToken(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}
