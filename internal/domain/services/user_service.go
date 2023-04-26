package services

import (
	"context"
	"github.com/elsyarif/go-auth-api/internal/domain/entities"
	"github.com/elsyarif/go-auth-api/internal/domain/repository"
	"github.com/elsyarif/go-auth-api/pkg/encryption"
	"github.com/elsyarif/go-auth-api/pkg/uid"
	"time"
)

type UserService struct {
	userRepo    repository.UserRepository
	idGenerator uid.NanoGenerator
	password    encryption.Password
}

func NewUserService(ur repository.UserRepository, uid uid.NanoGenerator, hash encryption.Password) UserService {
	return UserService{
		userRepo:    ur,
		idGenerator: uid,
		password:    hash,
	}
}

func (u *UserService) CreateUser(ctx context.Context, user entities.User) (*entities.User, error) {
	err := u.userRepo.VerifyAvailableUsername(ctx, user.Username)
	if err != nil {
		return nil, err
	}
	err = u.userRepo.VerifyAvailableEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	now := time.Now().Local()

	// Generate uuid
	user.Id = u.idGenerator.NanoId("user")
	// Hash password
	user.Password = u.password.Hash(user.Password)
	user.IsActive = false
	user.CreatedAt = now
	user.UpdatedAt = now
	err = u.userRepo.AddUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserService) GetUserById(ctx context.Context, id string) (*entities.User, error) {
	return u.userRepo.GetUserById(ctx, id)
}
