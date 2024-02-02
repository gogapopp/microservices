package handler

import (
	pb "github.com/gogapopp/microservices/chat-server/api/chat_v1"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedChatV1Server
}

func NewGRPCServer() *grpc.Server {
	grpcserver := grpc.NewServer()
	pb.RegisterChatV1Server(grpcserver, &GRPCServer{})
	return grpcserver
}
