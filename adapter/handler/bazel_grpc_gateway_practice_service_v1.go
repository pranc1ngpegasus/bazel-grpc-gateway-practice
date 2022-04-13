package handler

import (
	pb "github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/proto/bazel_grpc_gateway_practice/v1"
)

type bazelGrpcGatewayPracticeServiceV1 struct {
	pb.UnimplementedBazelGrpcGatewayPracticeServiceServer
}

func NewBazelGrpcGatewayPracticeServiceV1() pb.BazelGrpcGatewayPracticeServiceServer {
	return &bazelGrpcGatewayPracticeServiceV1{}
}
