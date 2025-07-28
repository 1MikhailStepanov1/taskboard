package services

import (
	"context"
	"log/slog"
	"taskboard/auth-service/internal/models"
	"taskboard/auth-service/internal/storage"
)

type UserServiceImpl struct {
	log     *slog.Logger
	storage storage.User
}

func NewUserService(
	log *slog.Logger,
	storage storage.User,
) *UserServiceImpl {
	return &UserServiceImpl{
		log:     log,
		storage: storage,
	}
}

type UserStorage interface {
	UserByEmail(ctx context.Context, email string) (*models.User, error)
	UserByShortName(ctx context.Context, shortName string) (*models.User, error)
	SaveUser(ctx context.Context, user *models.User) error
}

func (u *UserServiceImpl) Register(
	ctx context.Context,
	email string,
	password string,
	name string,
	surname string,
) (string, error) {
	// при регистрации приложуха сама придумывает уникальное короткое имя из name и surname
	// через bcrypt можно сделать формирование пароля сразу с солью и не парится с ее хранением
	return "", nil
}

func (u *UserServiceImpl) Login(
	ctx context.Context,
	email string,
	shortName string,
	password string,
) (string, error) {
	return "", nil
}
