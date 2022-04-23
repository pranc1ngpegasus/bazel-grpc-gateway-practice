package repository

import (
	"context"

	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/domain/model"
)

type (
	User interface {
		Create(ctx context.Context, user model.User) error
	}
)
