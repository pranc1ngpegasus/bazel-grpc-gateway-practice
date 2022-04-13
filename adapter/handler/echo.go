package handler

import (
	"context"

	pb "github.com/Pranc1ngPegasus/bazel-grpc-gateway-practice/proto/bazel_grpc_gateway_practice/v1"
)

var _ EchoProvider = (*echoProvider)(nil)

type (
	EchoProvider interface {
		Execute(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error)
	}

	echoProvider struct {
	}
)

func NewEchoProvider() EchoProvider {
	return &echoProvider{}
}

func (h *echoProvider) Execute(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Value: req.GetValue(),
	}, nil
}
