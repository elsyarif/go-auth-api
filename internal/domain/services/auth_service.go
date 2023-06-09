package services

import (
	"context"
	"errors"
	"github.com/elsyarif/go-auth-api/internal/domain/entities"
	"github.com/elsyarif/go-auth-api/internal/domain/repository"
	"github.com/elsyarif/go-auth-api/pkg/encryption"
	"net/mail"
)

type AuthService struct {
	authRepo repository.AuthRepository
	userRepo repository.UserRepository
	password encryption.Password
}

func NewAuthService(ar repository.AuthRepository, ur repository.UserRepository, ps encryption.Password) AuthService {
	return AuthService{
		authRepo: ar,
		userRepo: ur,
		password: ps,
	}
}

func (a *AuthService) CheckCredential(ctx context.Context, auth entities.AuthRequest) (string, error) {
	var passwordHash string
	var id string

	_, err := mail.ParseAddress(auth.Identity)
	if err != nil {
		id, passwordHash, err = a.userRepo.GetPasswordByUsername(ctx, auth.Identity)
		if err != nil {
			return "", errors.New("kredensial yang dimasukan salah")
		}
	} else {
		id, passwordHash, err = a.userRepo.GetPasswordByEmail(ctx, auth.Identity)
		if err != nil {
			return "", errors.New("kredensial yang dimasukan salah")
		}
	}

	err = a.password.ComparePassword(auth.Password, passwordHash)
	if err != nil {
		return "", errors.New("kredensial yang dimasukan salah")
	}

	return id, nil
}

func (a *AuthService) AddToken(ctx context.Context, request entities.RefreshTokenRequest) error {
	return a.authRepo.AddToken(ctx, request)
}

func (a *AuthService) CheckAvailabilityToken(ctx context.Context, request entities.RefreshTokenRequest) error {
	return a.authRepo.CheckAvailabilityToken(ctx, request)
}

func (a *AuthService) DeleteToken(ctx context.Context, request entities.RefreshTokenRequest) error {
	return a.authRepo.DeleteToken(ctx, request)
}
