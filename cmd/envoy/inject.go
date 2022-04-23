//go:generate go run github.com/google/wire/cmd/wire

package main

import (
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/configuration"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/server"

	"github.com/google/wire"
)

func initialize() server.HttpServer {
	wire.Build(
		configuration.Get,

		server.NewHttpServer,
	)

	return nil
}
