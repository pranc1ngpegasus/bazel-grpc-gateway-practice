//go:build wireinject
// +build wireinject

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
