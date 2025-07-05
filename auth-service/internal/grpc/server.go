package grpc

import (
	"context"
	"google.golang.org/grpc"
	authServicev1 "taskboard/auth-service/gen"
)

type AuthService interface {
	Login(
		ctx context.Context,
		email string,
		shortName string,
		password string,
	) (token string, err error)
	Register(
		ctx context.Context,
		email string,
		password string,
		name string,
		surname string,
		shortName string,
	) (userID string, err error)
	CheckPermission(
		ctx context.Context,
		userID string,
		action string,
	) (ok bool, err error)
}

type authServiceAPI struct {
	authServicev1.UnimplementedAuthServiceServer
	auth AuthService
}

func RegisterAuthService(gRPCServer *grpc.Server, auth AuthService) {
	authServicev1.RegisterAuthServiceServer(gRPCServer, &authServiceAPI{auth: auth})
}

// Слой контроллеров с валидацией входных данных
// Бизнес логика выполняется в сервисном слое auth-service/service

func (s *authServiceAPI) Login(
	ctx context.Context,
	req *authServicev1.LoginRequest,
) (*authServicev1.LoginResponse, error) {

}

func (s *authServiceAPI) Register(
	ctx context.Context,
	req *authServicev1.RegisterRequest,
) (*authServicev1.RegisterResponse, error) {

}

func (s *authServiceAPI) CheckPermission(
	ctx context.Context,
	req *authServicev1.CheckPermissionRequest,
) (*authServicev1.CheckPermissionResponse, error) {

}
