// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/configuration"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/server"
)

// Injectors from inject.go:

func initialize() server.HttpServer {
	config := configuration.Get()
	httpServer := server.NewHttpServer(config)
	return httpServer
}