package greedy

import "github.com/mig-elgt/tsp/vns"

type random struct {
	cluster *vns.Cluster
}

// NewRandom creates a reandom solution.
func NewRandom(cluster *vns.Cluster) *random {
	return &random{cluster}
}

// InitSolution creates a basic solution.
func (r *random) InitSolution() *vns.Solution {
	var customers []int
	for _, stop := range r.cluster.Stops {
		customers = append(customers, stop.StopID)
	}
	routeData := []vns.Route{
		{
			Start: 0,
			Size:  len(customers),
		},
	}
	return &vns.Solution{
		Cluster:   r.cluster,
		Customers: customers,
		Routes:    routeData,
	}
}
