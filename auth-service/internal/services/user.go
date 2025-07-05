package services

import (
	"context"
	"log/slog"
)

type UserService struct {
	log *slog.Logger
}

func NewUserService(
	log *slog.Logger,
) *UserService {
	return &UserService{
		log: log,
	}
}

func (u *UserService) Register(
	ctx context.Context,
) (string, error) {
	return "", nil
}

func (u *UserService) Login(
	ctx context.Context,
) (string, error) {
	return "", nil
}
