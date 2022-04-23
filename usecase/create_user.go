package usecase

import (
	"context"

	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/domain/model"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/domain/repository"
)

var _ CreateUser = (*createuser)(nil)

type (
	CreateUser interface {
		Do(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error)
	}

	createuser struct {
		userRepository repository.User
	}
)

func NewCreateUser(
	userRepository repository.User,
) CreateUser {
	return &createuser{
		userRepository: userRepository,
	}
}

type (
	CreateUserInput struct {
		Name string
	}

	CreateUserOutput struct {
		User model.User
	}
)

func (u *createuser) Do(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error) {
	userID := model.NewUUID()
	user := model.NewUser()
	user.SetID(userID)
	user.SetName(input.Name)

	if err := u.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	return &CreateUserOutput{
		User: user,
	}, nil
}
