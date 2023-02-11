package vns

import (
	"context"

	"github.com/mig-elgt/tsp/optimizer"
	pb "github.com/mig-elgt/tsp/proto/vns"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type vns struct {
	client pb.BasicVNSServiceClient
}

func (v *vns) Optimize(stops []*optimizer.Stop, matrix []float64) ([]*optimizer.Stop, float64, error) {

	var locations []*pb.Stop
	for _, s := range stops {
		locations = append(locations, &pb.Stop{
			ID:  int32(s.StopID),
			Lat: s.Location.Lat,
			Lng: s.Location.Lng,
		})
	}

	resp, err := v.client.Optimize(context.Background(), &pb.OptimizeRequest{Stops: locations, Matrix: matrix})
	if err != nil {
		logrus.Errorf("could not optimize route: %v", err)
		return nil, 0, errors.Wrapf(err, "could not optimize route")
	}

	var result []*optimizer.Stop
	for _, stop := range resp.Stops {
		result = append(result, stops[stop.ID-1])
	}

	return result, resp.TotalDistance, nil
}
