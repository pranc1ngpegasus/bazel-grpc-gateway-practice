package handler

import (
	"context"

	pb "github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/proto/user/v1"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/usecase"
)

var _ CreateUserProvider = (*createUserProvider)(nil)

type (
	CreateUserProvider interface {
		Do(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	}

	createUserProvider struct {
		createUserUsecase usecase.CreateUser
	}
)

func NewCreateUserProvider(
	createUserUsecase usecase.CreateUser,
) CreateUserProvider {
	return &createUserProvider{
		createUserUsecase: createUserUsecase,
	}
}

func (h *createUserProvider) Do(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	result, err := h.createUserUsecase.Do(ctx, usecase.CreateUserInput{
		Name: req.GetName(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Id:   result.User.ID(),
		Name: result.User.Name(),
	}, nil
}
