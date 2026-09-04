package server

import (
	"context"
	"example/first/proto"
	"example/first/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChatServer struct {
	proto.UnimplementedChatServiceServer
	Repo     *repository.ChatRepository
	UserRepo *repository.UserRepository
}

func (c *ChatServer) CreateChat(
	ctx context.Context,
	req *proto.CreateChatRequest,
) (*proto.CreateChatResponse, error) {
	if req.UserId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "Не указан обязательный")
	}

	user, err := c.UserRepo.GetById(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if user == nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	chatExtId, err := c.Repo.Create(user, req.Title)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &proto.CreateChatResponse{ExtId: chatExtId}, nil
}
