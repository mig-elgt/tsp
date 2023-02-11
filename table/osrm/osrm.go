package osrm

import "github.com/mig-elgt/tsp/table"

type osrm struct{}

// Fetch fetches a cost distance matrix for a set of locations
func (o *osrm) Fetch(locations []*table.Location) ([][]*table.Cost, error) {
	panic("not implemented") // TODO: Implement
}
