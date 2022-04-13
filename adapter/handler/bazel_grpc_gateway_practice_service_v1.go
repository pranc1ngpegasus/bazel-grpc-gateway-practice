package handler

import (
	"context"

	pb "github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/proto/bazel_grpc_gateway_practice/v1"
)

var _ pb.BazelGrpcGatewayPracticeServiceServer = (*bazelGrpcGatewayPracticeServiceV1)(nil)

type (
	BazelGrpcGatewayPracticeServiceV1 pb.BazelGrpcGatewayPracticeServiceServer

	bazelGrpcGatewayPracticeServiceV1 struct {
		pb.UnimplementedBazelGrpcGatewayPracticeServiceServer
		echoProvider EchoProvider
	}
)

func NewBazelGrpcGatewayPracticeServiceV1(
	echoProvider EchoProvider,
) BazelGrpcGatewayPracticeServiceV1 {
	return &bazelGrpcGatewayPracticeServiceV1{
		echoProvider: echoProvider,
	}
}

func (h *bazelGrpcGatewayPracticeServiceV1) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return h.echoProvider.Execute(ctx, req)
}
