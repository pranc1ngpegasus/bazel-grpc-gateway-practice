//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/configuration"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/handler"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/repository"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/server"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/infrastructure"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/usecase"

	"github.com/google/wire"
)

func initialize() server.GrpcServer {
	wire.Build(
		configuration.Get,

		infrastructure.NewRDBConnector,

		repository.NewUser,

		usecase.NewCreateUser,

		handler.NewCreateUserProvider,
		handler.NewUserServiceV1,
		handler.NewEchoProvider,
		handler.NewBazelGrpcGatewayPracticeServiceV1,

		server.NewGrpcServer,
	)

	return nil
}
