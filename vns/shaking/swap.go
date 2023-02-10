package shaking

import (
	"math/rand"
	"time"

	"github.com/mig-elgt/tsp/vns"
)

type swap struct{}

// NewSwap creates new swap shaker
func NewSwap() *swap {
	return &swap{}
}

// Shake performs a swap operation between two random stops(customers)
// and creates a new solution as result after the operation.
func (swap) Shake(so *vns.Solution, res *vns.Solution) {
	clusterCopy := make([]int, len(so.Customers))
	copy(clusterCopy, so.Customers)
	res.Cluster = so.Cluster
	res.Customers = clusterCopy
	res.Routes = so.Routes
	i := randInt(0, len(so.Customers)-1)
	j := randInt(0, len(so.Customers)-1)
	res.Customers[i], res.Customers[j] = res.Customers[j], res.Customers[i]
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
