package grpc

import (
	"context"

	pb "github.com/mig-elgt/tsp/proto/vns"
	"github.com/mig-elgt/tsp/vns"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	optimizer vns.Optimizer
}

func (h *handler) optimizeRoute(ctx context.Context, req *pb.OptimizeRequest) (*pb.OptimizeResponse, error) {
	stops, err := h.optimizer.Optimize(nil)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "stop locations empty")
	}
	var result []*pb.Stop
	for _, stop := range stops {
		result = append(result, &pb.Stop{
			ID:  int32(stop.StopID),
			Lat: stop.Location.Lat,
			Lng: stop.Location.Lng,
		})
	}
	return &pb.OptimizeResponse{Stops: result}, nil
}
