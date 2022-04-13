package main

import (
	"net"

	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/configuration"
	"github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/adapter/handler"
	pb "github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/proto/bazel_grpc_gateway_practice/v1"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func init() {
	configuration.Load()
}

func main() {
	env := configuration.Get()

	log.Info().Msgf("gRPC port: %s", env.Grpc.ServerPort)
	lis, err := net.Listen("tcp", ":"+env.Grpc.ServerPort)
	if err != nil {
		log.Error().Err(err)
		return
	}

	s := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             env.Grpc.EnforceMentPolicyMinTime,
				PermitWithoutStream: env.Grpc.EnforceMentPermitWithoutStream,
			},
		),
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				MaxConnectionIdle:     env.Grpc.MaxConnectionIdle,
				MaxConnectionAge:      env.Grpc.MaxConnectionAge,
				MaxConnectionAgeGrace: env.Grpc.MaxConnectionAgeGrace,
				Time:                  env.Grpc.Time,
				Timeout:               env.Grpc.Timeout,
			},
		),
	)
	bazelGrpcGatewayPracticeServiceV1 := handler.NewBazelGrpcGatewayPracticeServiceV1()
	pb.RegisterBazelGrpcGatewayPracticeServiceServer(s, bazelGrpcGatewayPracticeServiceV1)

	log.Info().Msg("start gRPC server.")
	if err = s.Serve(lis); err != nil {
		log.Error().Err(err)
		return
	}
}
