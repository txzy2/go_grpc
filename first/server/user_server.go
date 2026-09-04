package server

import (
	"context"
	"example/first/proto"
	"example/first/repository"
	"example/first/src"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	proto.UnimplementedUserServiceServer
	Repo *repository.UserRepository
}

func (s *UserServer) CreateUser(
	ctx context.Context,
	req *proto.CreateUserRequest,
) (*proto.UserResponse, error) {
	u, err := src.NewUser(req.Name)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.Repo.Create(u); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.UserResponse{
		Id:   int32(u.GetId()),
		Name: u.GetName(),
	}, nil
}

func (s *UserServer) GetUser(
	ctx context.Context,
	req *proto.GetUserRequest,
) (*proto.UserResponse, error) {
	if req.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "ID must be greater than 0")
	}

	user, err := s.Repo.GetById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user")
	}

	if user == nil {
		return nil, status.Error(codes.NotFound, "User is not found")
	}

	return &proto.UserResponse{
		Id:   int32(user.GetId()),
		Name: user.GetName(),
	}, nil
}
