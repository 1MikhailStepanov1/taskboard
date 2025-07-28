package storage

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"taskboard/auth-service/internal/models"
)

type User struct {
	*BaseStorage
}

func NewUserStorage(storage *BaseStorage) *User {
	return &User{storage}
}

func (u *User) UserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := u.pool.QueryRow(ctx, "SELECT * FROM users WHERE email = $1", email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Surname,
		&user.ShortName,
		&user.IsActive,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *User) UserByShortName(ctx context.Context, shortName string) (*models.User, error) {
	var user models.User
	err := u.pool.QueryRow(ctx, "SELECT * FROM users WHERE short_name = $1", shortName).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Surname,
		&user.ShortName,
		&user.IsActive,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *User) SaveUser(
	ctx context.Context,
	email string,
	password []byte,
	name string,
	surname string,
	shortName string,
) (*uuid.UUID, error) {
	var res uuid.UUID
	err := u.pool.QueryRow(
		ctx,
		"INSERT INTO users(id, email, password, name, surname, short_name) VALUES (gen_random_uuid(), $1, $2, $3, $4, $5) RETURNING id",
		email, password, name, surname, shortName,
	).Scan(&res)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return nil, ErrUserAlreadyExists
			}
		}
		return nil, err
	}
	return &res, nil
}
