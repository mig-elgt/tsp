package optimize

import "github.com/mig-elgt/tsp/vns"

type vnsOptimizer struct{}

func NewVNS() *vnsOptimizer {
	return &vnsOptimizer{}
}

func (vnsOptimizer) Optimize(cluster *vns.Cluster) ([]*vns.OptimalRoute, error) {
	panic("not implemented") // TODO: Implement
}
