package vns

import (
	"context"

	"github.com/mig-elgt/tsp/optimizer"
	pb "github.com/mig-elgt/tsp/proto/vns"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type vns struct {
	client pb.BasicVNSServiceClient
}

func New(addr string) (*vns, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "could not dial service connection")
	}
	return &vns{
		client: pb.NewBasicVNSServiceClient(conn),
	}, nil
}

// Optimize runs a route optimization for a set of stops and returns
// an optimal solution and the route total distance.
func (v *vns) Optimize(stops []*optimizer.Stop, matrix []float64) ([]*optimizer.Stop, float64, error) {
	// Preperate request stops
	var locations []*pb.Stop
	for i, s := range stops {
		s.StopID = i + 1
		locations = append(locations, &pb.Stop{
			ID:  int32(s.StopID),
			Lat: s.Location.Lat,
			Lng: s.Location.Lng,
		})
	}
	// Run optimization process
	resp, err := v.client.Optimize(context.Background(), &pb.OptimizeRequest{Stops: locations, Matrix: matrix})
	if err != nil {
		logrus.Errorf("could not optimize route: %v", err)
		return nil, 0, errors.Wrapf(err, "could not optimize route")
	}
	// Convert response stops to optimizer stops
	var result []*optimizer.Stop
	for _, stop := range resp.Stops {
		result = append(result, stops[stop.ID-1])
	}
	return result, resp.TotalDistance, nil
}
