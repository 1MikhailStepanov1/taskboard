package grpc

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"net"
	"taskboard/auth-service/internal/config"
	authService "taskboard/auth-service/internal/grpc"
	"taskboard/auth-service/internal/services"
)

// Структура конкретно того, что будет крутиться на grpc
type App struct {
	logger     *slog.Logger
	gRPCServer *grpc.Server
	config     *config.Config
}

func New(
	config *config.Config,
	logger *slog.Logger,
	userService *services.UserServiceImpl,
	roleService *services.RoleServiceImpl,
) *App {
	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived, logging.PayloadSent,
		),
	}

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			logger.Error("Recovered from panic", slog.Any("panic", p))

			return status.Errorf(codes.Internal, "internal error")
		}),
	}
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
		logging.UnaryServerInterceptor(
			func(l *slog.Logger) logging.Logger { // function for cast slog to logging.Logger
				return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
					l.Log(ctx, slog.Level(lvl), msg, fields...)
				})
			}(logger), loggingOpts...),
	))

	authService.RegisterAuthService(grpcServer, userService, roleService)

	return &App{
		logger:     logger,
		gRPCServer: grpcServer,
		config:     config,
	}
}

// Stub method to handle errors from Start() function
// TODO Search about error handling on app start
func (a *App) Run() {
	err := a.Start()
	if err != nil {
		panic(err)
	}
}
func (a *App) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", a.config.App.GRPCPort))
	if err != nil {
		return err
	}

	a.logger.Info("GRPC Server started")
	if err = a.gRPCServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (a *App) Stop() {
	a.logger.Info("Stopping GRPC Server")
	a.gRPCServer.GracefulStop()
}
