package optimize

import (
	"github.com/mig-elgt/tsp/vns"
	"github.com/mig-elgt/tsp/vns/greedy"
	"github.com/mig-elgt/tsp/vns/neighborhood"
	"github.com/mig-elgt/tsp/vns/shaking"
)

// vnsOptimizer describes an optimizer to run VNS algorithm
type vnsOptimizer struct{}

// NewNVS create new vnsOptimizer.
func NewVNS() *vnsOptimizer {
	return &vnsOptimizer{}
}

// Optimize computes a VNS algorithm given a Cluster pointer.
// It returns an Optimal Solution.
func (vnsOptimizer) Optimize(cluster *vns.Cluster) ([]*vns.OptimalRoute, error) {
	random := greedy.NewRandom(cluster)
	// Create Initial Solution using a greedy method.
	initSol := random.InitSolution()

	// Create Neighborhood Operators
	twoOptOpt := neighborhood.NewTwoOpt()
	// TODO: Implement Insertion method
	// TODO: Implement Reverse method
	// TODO: Implement 3-opt method

	// Create Shakers operators
	swapShaker := shaking.NewSwap()
	// TODO: Implement Insertion method
	// TODO: Implement Reverse method

	// Run VNS algorithm
	solution := vns.BasicVNS(initSol, twoOptOpt, swapShaker)

	// Create new stops with the optimal solution order.
	var stops []vns.Stop
	for _, customer := range solution.Customers {
		stops = append(stops, cluster.Stops[customer-1])
	}

	return []*vns.OptimalRoute{
		{
			Stops:   stops,
			Vehicle: solution.Cluster.Vehicles[0],
		},
	}, nil
}
