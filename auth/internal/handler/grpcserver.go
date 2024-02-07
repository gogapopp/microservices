package handler

import (
	pb "github.com/gogapopp/microservices/auth/api/auth_v1"
	"github.com/gogapopp/microservices/auth/internal/repository"
	"github.com/gogapopp/microservices/auth/internal/repository/sqlite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	repository repository.Repository
	pb.UnimplementedAuthV1Server
}

func NewGRPCServer(repository *sqlite.Repository) *grpc.Server {
	grpcserver := grpc.NewServer()
	reflection.Register(grpcserver)
	pb.RegisterAuthV1Server(grpcserver, &GRPCServer{repository: repository})
	return grpcserver
}
