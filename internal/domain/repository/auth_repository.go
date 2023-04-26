package repository

import (
	"context"
	"github.com/elsyarif/go-auth-api/internal/domain/entities"
)

type AuthRepository interface {
	AddToken(ctx context.Context, request entities.RefreshTokenRequest) error
	CheckAvailabilityToken(ctx context.Context, request entities.RefreshTokenRequest) error
	DeleteToken(ctx context.Context, request entities.RefreshTokenRequest) error
}
