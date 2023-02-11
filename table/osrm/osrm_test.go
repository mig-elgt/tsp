package osrm

import (
	"testing"

	"github.com/mig-elgt/tsp/table"
)

func Test_osrmFetch(t *testing.T) {}

type TableServiceMock struct {
	FetchFn func(locations []*table.Location) ([][]*table.Cost, error)
}

// Fetch fetches a cost distance matrix for a set of locations
func (t *TableServiceMock) Fetch(locations []*table.Location) ([][]*table.Cost, error) {
	return t.FetchFn(locations)
}
