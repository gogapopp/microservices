package handler

import (
	pb "github.com/gogapopp/microservices/chat-server/api/chat_v1"
	"github.com/gogapopp/microservices/chat-server/internal/repository"
	"github.com/gogapopp/microservices/chat-server/internal/repository/sqlite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	repository repository.Repository
	pb.UnimplementedChatV1Server
}

func NewGRPCServer(repository *sqlite.Repository) *grpc.Server {
	grpcserver := grpc.NewServer()
	reflection.Register(grpcserver)
	pb.RegisterChatV1Server(grpcserver, &GRPCServer{repository: repository})
	return grpcserver
}
