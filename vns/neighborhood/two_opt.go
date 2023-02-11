package neighborhood

import (
	"math/rand"
	"time"

	"github.com/mig-elgt/tsp/vns"
)

type twoOpt struct {
	rand Randomer
}

func NewTwoOpt() *twoOpt {
	return &twoOpt{
		rand: &intRandom{},
	}
}

// BestImprovement chooses the best solution from N('s) where all possible moves are performed
func (t *twoOpt) BestImprovement(s *vns.Solution) *vns.Solution {
	moves := t.generateMoves(s)
	bestMove := EvaluateMovesAndReturnTheBestOne(moves, s, 1)
	newSol := t.applyMove(bestMove.(*twoOptMove), s)
	newSol.Cluster = s.Cluster
	return newSol
}

// The first improvement strategy tries to avoid the time complexity of exploring
func (t *twoOpt) FirstImprovement(s *vns.Solution) *vns.Solution {
	panic("not implemented") // TODO: Implement
}

type twoOptMove struct {
	rA, i, j int
}

type Randomer interface {
	Int(min, max int) int
}

type intRandom struct{}

func (intRandom) Int(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func (t twoOpt) generateMoves(s *vns.Solution) []SolutionMover {
	idx := t.rand.Int(0, len(s.Routes)-1)
	moves := []SolutionMover{}
	for i := s.Routes[idx].Start; i < s.Routes[idx].Start+s.Routes[idx].Size-1; i++ {
		for j := i + 1; j < s.Routes[idx].Start+s.Routes[idx].Size; j++ {
			moves = append(moves, &twoOptMove{idx, i, j})
		}
	}
	return moves
}

func (to twoOptMove) Evaluate(s *vns.Solution, k int) float64 {
	// Create Routes
	ra := s.Routes[to.rA]
	routeA := make([]int, ra.Size)
	copy(routeA, s.Customers[ra.Start:ra.Start+ra.Size])
	// Apply Move
	for k, l := to.i-ra.Start, to.j-ra.Start; k < l; k, l = k+1, l-1 {
		routeA[k], routeA[l] = routeA[l], routeA[k]
	}
	// Compute Route Cost
	newCost := vns.CalculateVehicleRouteCost(to.rA, routeA, s.Cluster)
	return (s.Fitness() - ra.Cost) + newCost
}

func (to twoOpt) applyMove(tom *twoOptMove, s *vns.Solution) *vns.Solution {
	// Create Routes
	customers := make([]int, len(s.Customers))
	copy(customers, s.Customers)
	routes := make([]vns.Route, len(s.Routes))
	copy(routes, s.Routes)
	ra := &routes[tom.rA]
	routeA := customers[ra.Start : ra.Start+ra.Size]
	// Apply Move
	for k, l := tom.i-ra.Start, tom.j-ra.Start; k < l; k, l = k+1, l-1 {
		routeA[k], routeA[l] = routeA[l], routeA[k]
	}
	// Compute Route Costs
	ra.Cost = vns.CalculateVehicleRouteCost(tom.rA, routeA, s.Cluster)
	// Create Solution and calculate fitness value
	ns := &vns.Solution{
		Customers:    customers,
		Routes:       routes,
		FitnessValue: 1 / ra.Cost,
	}
	return ns

}
