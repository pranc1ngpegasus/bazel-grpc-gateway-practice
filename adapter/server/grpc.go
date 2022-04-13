package server

import (
	"net"

	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/configuration"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/handler"
	pb "github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/proto/bazel_grpc_gateway_practice/v1"

	"github.com/philip-bui/grpc-zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type (
	GrpcServer interface {
		Serve() error
	}

	grpcServer struct {
		server                            *grpc.Server
		config                            configuration.Config
		bazelGrpcGatewayPracticeServiceV1 handler.BazelGrpcGatewayPracticeServiceV1
	}
)

func NewGrpcServer(
	config configuration.Config,
	bazelGrpcGatewayPracticeServiceV1 handler.BazelGrpcGatewayPracticeServiceV1,
) GrpcServer {
	server := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             config.Grpc.EnforceMentPolicyMinTime,
				PermitWithoutStream: config.Grpc.EnforceMentPermitWithoutStream,
			},
		),
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				MaxConnectionIdle:     config.Grpc.MaxConnectionIdle,
				MaxConnectionAge:      config.Grpc.MaxConnectionAge,
				MaxConnectionAgeGrace: config.Grpc.MaxConnectionAgeGrace,
				Time:                  config.Grpc.Time,
				Timeout:               config.Grpc.Timeout,
			},
		),
		zerolog.UnaryInterceptor(),
	)

	pb.RegisterBazelGrpcGatewayPracticeServiceServer(server, bazelGrpcGatewayPracticeServiceV1)

	return &grpcServer{
		server: server,
		config: config,
	}
}

func (s *grpcServer) Serve() error {
	log.Info().Msgf("gRPC port: %s", s.config.Grpc.ServerPort)
	lis, err := net.Listen("tcp", ":"+s.config.Grpc.ServerPort)
	if err != nil {
		log.Error().Err(err).Caller().Stack()
		return err
	}

	log.Info().Msg("start gRPC server.")
	if err = s.server.Serve(lis); err != nil {
		log.Error().Err(err).Caller().Stack()
		return err
	}

	return nil
}
