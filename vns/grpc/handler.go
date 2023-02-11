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
	stops, err := h.optimizer.Optimize(h.createCluster(req))
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

func (h *handler) createCluster(req *pb.OptimizeRequest) *vns.Cluster {
	var stops []vns.Stop
	// Create stop list
	for _, stop := range req.Stops {
		stops = append(stops, vns.Stop{
			StopID: int(stop.ID),
			Location: &vns.Location{
				Lat: stop.Lat,
				Lng: stop.Lng,
			},
		})
	}
	// Create Cost Matrix
	lenStops := len(stops) + 1
	var matrix vns.CostMatrix
	for i := 0; i < len(req.Matrix); i += lenStops {
		row := []vns.TravelCost{}
		for j := 0; j < lenStops; j++ {
			row = append(row, vns.TravelCost{
				Distance: req.Matrix[j+i],
			})
		}
		matrix = append(matrix, row)
	}
	return &vns.Cluster{Stops: stops, CostMatrix: matrix}
}
