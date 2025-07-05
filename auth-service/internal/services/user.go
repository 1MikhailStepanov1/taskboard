package services

import (
	"context"
	"log/slog"
)

type UserServiceImpl struct {
	log *slog.Logger
}

func NewUserService(
	log *slog.Logger,
) *UserServiceImpl {
	return &UserServiceImpl{
		log: log,
	}
}

func (u *UserServiceImpl) Register(
	ctx context.Context,
	email string,
	password string,
	name string,
	surname string,
	shortName string,
) (string, error) {
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
