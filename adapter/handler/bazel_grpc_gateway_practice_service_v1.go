package handler

import (
	pb "github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/proto/bazel_grpc_gateway_practice/v1"
)

var _ pb.BazelGrpcGatewayPracticeServiceServer = (*bazelGrpcGatewayPracticeServiceV1)(nil)

type (
	BazelGrpcGatewayPracticeServiceV1 pb.BazelGrpcGatewayPracticeServiceServer

	bazelGrpcGatewayPracticeServiceV1 struct {
		pb.UnimplementedBazelGrpcGatewayPracticeServiceServer
	}
)

func NewBazelGrpcGatewayPracticeServiceV1() BazelGrpcGatewayPracticeServiceV1 {
	return &bazelGrpcGatewayPracticeServiceV1{}
}
