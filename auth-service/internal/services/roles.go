package services

import (
	"context"
	"log/slog"
)

type RoleService struct {
	log *slog.Logger
}

func NewRoleService(
	log *slog.Logger,
) *RoleService {
	return &RoleService{log: log}
}

func (r *RoleService) CheckPermission(
	ctx context.Context,
) (bool, error) {
	return true, nil
}
