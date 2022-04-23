package handler

import (
	"context"

	pb "github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/proto/user/v1"
)

type (
	UserServiceV1 pb.UserServiceServer

	userServiceV1 struct {
		pb.UnimplementedUserServiceServer
		createUserProvider CreateUserProvider
	}
)

func NewUserServiceV1(
	createUserProvider CreateUserProvider,
) UserServiceV1 {
	return &userServiceV1{
		createUserProvider: createUserProvider,
	}
}

func (h *userServiceV1) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return h.createUserProvider.Do(ctx, req)
}
