package handler

import (
	pb "github.com/gogapopp/microservices/auth/api/auth_v1"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedAuthV1Server
}

func NewGRPCServer() *grpc.Server {
	grpcserver := grpc.NewServer()
	pb.RegisterAuthV1Server(grpcserver, &GRPCServer{})
	return grpcserver
}
