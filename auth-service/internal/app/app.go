package app

import "google.golang.org/grpc"

type App struct {
	GRPCServer *grpc.Server
}
