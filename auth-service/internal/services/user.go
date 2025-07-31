package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"taskboard/auth-service/internal/config"
	"taskboard/auth-service/internal/lib/jwt"
	"taskboard/auth-service/internal/models"
	"taskboard/auth-service/internal/storage"
)

type UserServiceImpl struct {
	log     *slog.Logger
	storage storage.User
	config  config.SecurityConfig
}

func NewUserService(
	log *slog.Logger,
	storage storage.User,
	config config.SecurityConfig,
) *UserServiceImpl {
	return &UserServiceImpl{
		log:     log,
		storage: storage,
		config:  config,
	}
}

type UserStorage interface {
	UserByEmail(ctx context.Context, email string) (*models.User, error)
	UserByShortName(ctx context.Context, shortName string) (*models.User, error)
	SaveUser(ctx context.Context,
		email string,
		password []byte,
		name string,
		surname string,
		shortName string,
	) (*uuid.UUID, error)
}

func (u *UserServiceImpl) Register(
	ctx context.Context,
	email string,
	password string,
	name string,
	surname string,
	shortName string,
) (string, error) {
	u.log.Info("Registering user", "email", email, "name", name, "surname", surname)
	// Проверка на наличие такого пользователя
	_, err := u.storage.UserByEmail(ctx, email)
	if err == nil {
		return "", errors.New("user with this email already exists")
	}

	// Хеширование пароля
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// Создание пользователя
	id, err := u.storage.SaveUser(ctx, email, passwordHash, name, surname, shortName)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (u *UserServiceImpl) Login(
	ctx context.Context,
	email string,
	shortName string,
	password string,
) (string, error) {
	var (
		user *models.User
		err  error
	)
	if email != "" {
		user, err = u.storage.UserByEmail(ctx, email)
		if err != nil {
			return "", err
		}
	} else if shortName != "" {
		user, err = u.storage.UserByShortName(ctx, shortName)
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("not enough data to identify user")
	}
	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		return "", errors.New("password incorrect") // TODO сделать нормальную обработку ошибок
	}
	token, err := jwt.GenerateToken(*user, u.config.JWTSecret, u.config.JWTDuration)
	if err != nil {
		return "", err
	}
	return token, nil
}
