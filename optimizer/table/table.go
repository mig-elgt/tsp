package table

import (
	"context"

	"github.com/mig-elgt/tsp/optimizer"
	pb "github.com/mig-elgt/tsp/proto/table"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type table struct {
	client pb.TableServiceClient
}

func (t *table) GetDistanceMatrix(stops []*optimizer.Stop) ([]float64, error) {
	resp, err := t.client.Fetch(context.Background(), &pb.FetchRequest{})
	if err != nil {
		logrus.Errorf("could not fetch distance matrix: %v", err)
		return nil, errors.Wrapf(err, "could not fetch distance matrix")
	}
	return resp.Matrix, nil
}
