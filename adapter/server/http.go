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
		server *runtime.ServeMux
		config configuration.Config
	}
)

func NewHttpServer(
	config configuration.Config,
) HttpServer {
	return &httpServer{
		server: runtime.NewServeMux(),
		config: config,
	}
}

func (s *httpServer) Serve() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	}

	log.Info().Msgf("listen gRPC port: %s", s.config.Grpc.ServerPort)
	if err := pb.RegisterBazelGrpcGatewayPracticeServiceHandlerFromEndpoint(ctx, s.server, ":"+s.config.Grpc.ServerPort, opts); err != nil {
		log.Error().Stack().Caller().Err(err)
		panic(err)
	}

	log.Info().Msgf("HTTP port: %s", s.config.Http.ServerPort)
	if err := http.ListenAndServe(":"+s.config.Http.ServerPort, s.server); err != nil {
		log.Error().Stack().Caller().Err(err)
		return err
	}

	return nil
}
