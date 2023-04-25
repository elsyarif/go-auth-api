package repositories

import (
	"context"
	"errors"
	"github.com/elsyarif/go-auth-api/internal/domain/entities"
	"github.com/elsyarif/go-auth-api/internal/domain/repository"
	"github.com/elsyarif/go-auth-api/pkg/helper/log"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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
	query := "INSERT INTO users values ($1, $2, $3, $4, $5, $6, $7, $8)"

	tx, err := u.DB.Beginx()
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, query, user.Id, user.Name, user.Username, user.Email, user.Password, user.IsActive, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		log.Error("exec add user error", logrus.Fields{"error": err})
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
	query := "SELECT username FROM users WHERE username = $1"
	user := entities.User{}

	tx, err := u.DB.Beginx()
	if err != nil {
		return err
	}

	_ = tx.GetContext(ctx, &user, query, username)

	if user.Username != "" {
		log.Warn("username already exists with", logrus.Fields{
			"username": user.Username,
		})
		return errors.New("username already exists")
	}
	return nil
}

func (u *UserRepositoryPostgres) VerifyAvailableEmail(ctx context.Context, email string) error {
	query := "SELECT email FROM users WHERE email = $1"
	user := entities.User{}

	tx, err := u.DB.Beginx()
	if err != nil {
		return err
	}

	_ = tx.GetContext(ctx, &user, query, email)

	if user.Email != "" {
		log.Warn("email already exists with", logrus.Fields{
			"email": user.Email,
		})
		return errors.New("email already exists")
	}
	return nil
}

func (u *UserRepositoryPostgres) GetPasswordByUsername(ctx context.Context, username string) (string, error) {
	query := "SELECT password FROM users WHERE username = $1"
	user := entities.User{}

	tx, err := u.DB.Beginx()
	if err != nil {
		log.Error("database error", logrus.Fields{"with": err})
		return "", err
	}

	err = tx.GetContext(ctx, &user, query, username)
	if err != nil {
		log.Warn("user tidak ditemukan", logrus.Fields{"username": username})
		return "", err
	}

	return user.Password, nil
}

func (u *UserRepositoryPostgres) GetIdByUsername(ctx context.Context, username string) (string, error) {
	query := "SELECT id FROM users WHERE username = $1"
	user := entities.User{}

	tx, err := u.DB.Beginx()
	if err != nil {
		log.Error("database error", logrus.Fields{"with": err})
		return "", err
	}

	err = tx.GetContext(ctx, &user, query, username)
	if err != nil {
		log.Warn("user tidak ditemukan", logrus.Fields{"username": username})
		return "", err
	}

	return user.Id, nil
}
