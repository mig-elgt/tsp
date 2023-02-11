package grpc

import (
	"context"

	pb "github.com/mig-elgt/tsp/proto/table"
	"github.com/mig-elgt/tsp/table"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	table table.TableService
}

func (h *handler) fetchDistanceMatrix(ctx context.Context, req *pb.FetchRequest) (*pb.FetchResponse, error) {
	if len(req.Stops) == 0 {
		return nil, status.Error(codes.InvalidArgument, "stop locations is empty")
	}
	// Convert stops locations to table locations
	var locations []*table.Location
	for _, stop := range req.Stops {
		locations = append(locations, &table.Location{
			Lat: stop.Lat,
			Lng: stop.Lng,
		})
	}
	// Fetch distance matrix from table service
	matrix, err := h.table.Fetch(locations)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not get distance matrix: %v", err)
	}
	// Convert distance matrix and stores distance values
	var distances []float64
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			distances = append(distances, matrix[i][j].Distance)
		}
	}

	return &pb.FetchResponse{Matrix: distances}, nil
}
