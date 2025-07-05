package grpc

import (
	"context"
	"google.golang.org/grpc"
	authServicev1 "taskboard/auth-service/gen"
)

type UserService interface {
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
}

type RoleService interface {
	CheckPermission(
		ctx context.Context,
		userID string,
		action string,
	) (ok bool, err error)
}

type authServiceAPI struct {
	authServicev1.UnimplementedAuthServiceServer
	user UserService
	role RoleService
}

func RegisterAuthService(gRPCServer *grpc.Server, user UserService, role RoleService) {
	authServicev1.RegisterAuthServiceServer(gRPCServer, &authServiceAPI{user: user, role: role})
}

// Слой контроллеров с валидацией входных данных
// Бизнес логика выполняется в сервисном слое auth-service/service

func (s *authServiceAPI) Login(
	ctx context.Context,
	req *authServicev1.LoginRequest,
) (*authServicev1.LoginResponse, error) {
	token, err := s.user.Login(ctx, req.Email, req.ShortName, req.Password)
	if err != nil {
	}
	return &authServicev1.LoginResponse{Token: token}, nil
}

func (s *authServiceAPI) Register(
	ctx context.Context,
	req *authServicev1.RegisterRequest,
) (*authServicev1.RegisterResponse, error) {
	userID, err := s.user.Register(ctx, req.Email, req.Password, req.Name, req.Surname, req.ShortName)
	if err != nil {
	}
	return &authServicev1.RegisterResponse{UserId: userID}, nil
}

func (s *authServiceAPI) CheckPermission(
	ctx context.Context,
	req *authServicev1.CheckPermissionRequest,
) (*authServicev1.CheckPermissionResponse, error) {
	ok, err := s.role.CheckPermission(ctx, req.ShortName, req.Action)
	if err != nil {
	}
	return &authServicev1.CheckPermissionResponse{Ok: ok}, err
}
