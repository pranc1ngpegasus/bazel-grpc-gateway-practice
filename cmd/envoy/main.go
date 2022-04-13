package main

import (
	"context"
	"net/http"

	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/configuration"
	pb "github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/proto/bazel_grpc_gateway_practice/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	configuration.Load()
}

func main() {
	env := configuration.Get()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	}

	if err := pb.RegisterBazelGrpcGatewayPracticeServiceHandlerFromEndpoint(ctx, mux, ":"+env.Grpc.ServerPort, opts); err != nil {
		log.Error().Err(err)
		return
	}

	if err := http.ListenAndServe(":"+env.Http.ServerPort, mux); err != nil {
		log.Error().Err(err)
		return
	}
}
