package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"taskboard/auth-service/internal/app"
	"taskboard/auth-service/internal/config"
)

const (
	localLogsLevel      = "LOCAL"
	productionLogsLevel = "PROD"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case localLogsLevel:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case productionLogsLevel:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func main() {
	cfg := config.New()

	logger := setupLogger(cfg.App.LogLevel)

	application := app.New(*cfg, *logger)

	go func() {
		application.GRPCApp.Run()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	application.GRPCApp.Stop()
	logger.Info("GRPC gracefully stopped")
}
