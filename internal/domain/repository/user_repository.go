package repository

import (
	"context"
	"github.com/elsyarif/go-auth-api/internal/domain/entities"
)

type UserRepository interface {
	AddUser(ctx context.Context, user entities.User) error
	VerifyAvailableUsername(ctx context.Context, username string) error
	VerifyAvailableEmail(ctx context.Context, email string) error
	GetPasswordByUsername(ctx context.Context, username string) (string, error)
	GetIdByUsername(ctx context.Context, username string) (string, error)
}
