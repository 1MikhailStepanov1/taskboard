package services

import (
	"context"
	"log/slog"
	"taskboard/auth-service/internal/storage"
)

type RoleServiceImpl struct {
	log         *slog.Logger
	roleStorage *storage.Roles
}

func NewRoleService(
	log *slog.Logger,
	roleStorage *storage.Roles,
) *RoleServiceImpl {
	return &RoleServiceImpl{log: log, roleStorage: roleStorage}
}

func (r *RoleServiceImpl) CheckPermission(
	ctx context.Context,
	userID string,
	action string,
) (bool, error) {
	return true, nil
}
