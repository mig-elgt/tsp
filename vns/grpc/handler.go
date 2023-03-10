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
	cluster := h.createCluster(req)
	stops, err := h.optimizer.Optimize(cluster)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// Create stops response and calculate the total distance
	var totalDistance float64
	var result []*pb.Stop
	for i := 0; i < len(stops); i++ {
		stop := stops[i]
		result = append(result, &pb.Stop{
			ID:  int32(stop.StopID),
			Lat: stop.Location.Lat,
			Lng: stop.Location.Lng,
		})
		if i+1 < len(stops) {
			totalDistance += cluster.CostMatrix[i][i+1].Distance
		}
	}
	totalDistance += cluster.CostMatrix[stops[len(stops)-1].StopID-1][0].Distance

	return &pb.OptimizeResponse{Stops: result, TotalDistance: totalDistance}, nil
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
	lenStops := len(stops)
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
