package repository

import (
	"context"

	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/domain/model"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/domain/repository"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/infrastructure"
)

var _ repository.User = (*userDAO)(nil)

type (
	userDAO struct {
		rdbConnector infrastructure.RDBConnector
	}
)

func NewUser(
	rdbConnector infrastructure.RDBConnector,
) repository.User {
	return &userDAO{
		rdbConnector: rdbConnector,
	}
}

func (r *userDAO) Create(ctx context.Context, user model.User) error {
	if _, err := r.rdbConnector.GetDB().
		User.
		Create().
		SetID(
			user.ID(),
		).
		SetName(
			user.Name(),
		).
		Save(ctx); err != nil {
		return err
	}

	return nil
}
