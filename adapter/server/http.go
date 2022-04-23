package server

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

type (
	HttpServer interface {
		Serve() error
	}

	httpServer struct {
		server http.Handler
		config configuration.Config
	}
)

func NewHttpServer(
	config configuration.Config,
) HttpServer {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	}

	if err := pb.RegisterBazelGrpcGatewayPracticeServiceHandlerFromEndpoint(ctx, mux, ":"+config.Grpc.ServerPort, opts); err != nil {
		log.Error().Stack().Caller().Err(err)
		panic(err)
	}

	return &httpServer{
		server: mux,
		config: config,
	}
}

func (s *httpServer) Serve() error {
	log.Info().Msgf("HTTP port: %s", s.config.Http.ServerPort)
	if err := http.ListenAndServe(":"+s.config.Http.ServerPort, s.server); err != nil {
		log.Error().Stack().Caller().Err(err)
		return err
	}

	return nil
}
