package handler

import (
	"time"

	pb "github.com/gogapopp/microservices/chat-server/api/chat_v1"
	"github.com/gogapopp/microservices/chat-server/internal/models"

	"context"
)

func (g *GRPCServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	err := g.repository.Create(ctx, models.Message{
		From:      "german",
		Text:      "german",
		Timestamp: time.Now(),
	})
	if err == nil {
		return nil, nil
	}
	return &pb.CreateResponse{
		Id: 1,
	}, nil
}

func (g *GRPCServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return &pb.DeleteResponse{}, nil
}

func (g *GRPCServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	return &pb.SendMessageResponse{}, nil
}
