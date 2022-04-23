package main

import (
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/configuration"

	"github.com/rs/zerolog/log"
)

func init() {
	configuration.Load()
}

func main() {
	server := initialize()

	if err := server.Serve(); err != nil {
		log.Error().Err(err).Caller().Stack()
		return
	}
}
