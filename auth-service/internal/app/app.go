package app

import (
	"log/slog"
	"taskboard/auth-service/internal/app/grpc"
	"taskboard/auth-service/internal/config"
	"taskboard/auth-service/internal/services"
	"taskboard/auth-service/internal/storage"
)

type App struct {
	GRPCApp *grpc.App
}

// Общая собирательная структура приложения
func New(config config.Config, logger slog.Logger) *App {
	connPool, err := storage.NewConnectionPool(config.DB)
	if err != nil {
		panic(err)
	}
	userStorage := storage.NewUser(connPool)
	roleStorage := storage.NewRoles(connPool)

	userService := services.NewUserService(&logger, *userStorage, config.Security)
	roleService := services.NewRoleService(&logger, roleStorage)

	grpcApp := grpc.New(&config, &logger, userService, roleService)
	return &App{GRPCApp: grpcApp}
}
