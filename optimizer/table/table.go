package table

import (
	"context"

	"github.com/mig-elgt/tsp/optimizer"
	pb "github.com/mig-elgt/tsp/proto/table"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type table struct {
	client pb.TableServiceClient
}

func New(addr string) (*table, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "could not dial service connection")
	}
	return &table{
		client: pb.NewTableServiceClient(conn),
	}, nil
}

func (t *table) GetDistanceMatrix(stops []*optimizer.Stop) ([]float64, error) {
	var stopLocations []*pb.Stop
	for _, s := range stops {
		stopLocations = append(stopLocations, &pb.Stop{
			Lat: s.Location.Lat,
			Lng: s.Location.Lng,
		})
	}
	resp, err := t.client.Fetch(context.Background(), &pb.FetchRequest{
		Stops: stopLocations,
	})
	if err != nil {
		logrus.Errorf("could not fetch distance matrix: %v", err)
		return nil, errors.Wrapf(err, "could not fetch distance matrix")
	}
	return resp.Matrix, nil
}
