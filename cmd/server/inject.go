//go:build wireinject

package main

import (
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/configuration"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/handler"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/server"

	"github.com/google/wire"
)

func initialize() server.GrpcServer {
	wire.Build(
		configuration.Get,

		handler.NewEchoProvider,
		handler.NewBazelGrpcGatewayPracticeServiceV1,

		server.NewGrpcServer,
	)

	return nil
}
