package grpc

import (
	"context"

	pb "github.com/mig-elgt/tsp/proto/vns"
)

type handler struct{}

func (h *handler) optimizeRoute(_ context.Context, _ *pb.OptimizeRequest) (*pb.OptimizeResponse, error) {
	panic("not implemented") // TODO: Implement
}
