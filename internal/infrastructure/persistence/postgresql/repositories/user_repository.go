package repositories

import (
	"auth-hexa/internal/domain/entities"
	"auth-hexa/internal/domain/repository"
	"auth-hexa/pkg/helper/logger"
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type UserRepositoryPostgres struct {
	DB *sqlx.DB
}

func NewUserRepositoryPostgres(db *sqlx.DB) repository.UserRepository {
	return &UserRepositoryPostgres{
		DB: db,
	}
}

func (u *UserRepositoryPostgres) AddUser(ctx context.Context, user entities.User) error {
	query := "INSERT INTO user_test values ($1, $2, $3, $4, $5, $6, $7)"

	tx, err := u.DB.Beginx()
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, query, user.Id, user.Name, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if row, err := result.RowsAffected(); err == nil && row > 0 {
		_ = tx.Commit()
		return nil
	}

	return err
}

func (u *UserRepositoryPostgres) VerifyAvailableUsername(ctx context.Context, username string) error {
	query := "SELECT username FROM user_test WHERE username = $1"
	user := entities.User{}

	tx, err := u.DB.Beginx()
	if err != nil {
		return err
	}

	_ = tx.GetContext(ctx, &user, query, username)

	if user.Username != "" {
		logger.Error("username already exists", zap.Error(errors.New("VerifyAvailableUsername")))
		return errors.New("username already exists")
	}
	return nil
}

func (u *UserRepositoryPostgres) GetPasswordByUsername(ctx context.Context, username string) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepositoryPostgres) GetIdByUsername(ctx context.Context, username string) error {
	//TODO implement me
	panic("implement me")
}
