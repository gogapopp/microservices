package handler

import (
	pb "github.com/gogapopp/microservices/auth/api/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	pb.UnimplementedAuthV1Server
}

func NewGRPCServer() *grpc.Server {
	grpcserver := grpc.NewServer()
	reflection.Register(grpcserver)
	pb.RegisterAuthV1Server(grpcserver, &GRPCServer{})
	return grpcserver
}
