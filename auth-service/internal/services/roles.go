package services

import (
	"context"
	"log/slog"
)

type RoleServiceImpl struct {
	log *slog.Logger
}

func NewRoleService(
	log *slog.Logger,
) *RoleServiceImpl {
	return &RoleServiceImpl{log: log}
}

func (r *RoleServiceImpl) CheckPermission(
	ctx context.Context,
	userID string,
	action string,
) (bool, error) {
	return true, nil
}
