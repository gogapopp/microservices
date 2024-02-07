package handler

import (
	"time"

	pb "github.com/gogapopp/microservices/auth/api/auth_v1"
	"github.com/gogapopp/microservices/auth/models"
	"google.golang.org/protobuf/types/known/timestamppb"

	"context"
)

func (g *GRPCServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	err := g.repository.Create(ctx, models.User{
		Userame:  "german",
		Email:    "german@email.com",
		Password: "german",
		Role:     "1",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	})
	if err == nil {
		return nil, nil
	}
	return &pb.CreateResponse{
		Id: 1,
	}, nil
}

func (g *GRPCServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{
		Id:        1,
		Name:      "123",
		Email:     "123",
		Role:      1,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}, nil
}

func (g *GRPCServer) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	return &pb.UpdateResponse{}, nil
}

func (g *GRPCServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return &pb.DeleteResponse{}, nil
}
