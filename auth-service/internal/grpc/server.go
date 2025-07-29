package grpc

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	authServicev1 "taskboard/auth-service/gen"
	"taskboard/auth-service/internal/storage"
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
	if req.Email == "" && req.ShortName == "" {
		return nil, status.Error(codes.InvalidArgument, "email or short name required")
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password required")
	}
	token, err := s.user.Login(ctx, req.Email, req.ShortName, req.Password)
	if err != nil {
	}
	return &authServicev1.LoginResponse{Token: token}, nil
}

func (s *authServiceAPI) Register(
	ctx context.Context,
	req *authServicev1.RegisterRequest,
) (*authServicev1.RegisterResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	if req.Surname == "" {
		return nil, status.Error(codes.InvalidArgument, "surname is required")
	}
	if req.ShortName == "" {
		return nil, status.Error(codes.InvalidArgument, "short name is required")
	}
	userID, err := s.user.Register(ctx, req.Email, req.Password, req.Name, req.Surname, req.ShortName)
	if err != nil {
		if errors.Is(err, storage.ErrUserAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Error(codes.Internal, "failed to register user")
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
