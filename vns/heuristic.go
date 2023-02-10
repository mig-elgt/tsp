// vns describes an algorithm based in metaheuristic that starts searching with a single solution.
// They could be see as walks through neihborhoods on search trajectories
// through the search space of the problem at hand.
// The walks are performing by iterative procedures that move from the current
// solution to another one in the search space.

package vns

// Solution describes the representation to VRP problem
type Solution struct {
	// cluster holds the optimization stops, vehicles and distance matrix information
	Cluster *Cluster
	// customers representation as a sequence vector
	Customers []int
	// Routes describes a set of vehicles routes.
	Routes []Route
	// Fitness values describes the solution quality
	FitnessValue float32
}

// Route stores the vehicle routes information
type Route struct {
	Start int
	Size  int
	Cost  float32
}

// HeuristicConstructor describes the behavior to generate a greedy solutions.
// It generates a solution from scratch in a step-by-step manner.
type HeuristicConstructor interface {
	InitSolution() *Solution
}

// NeighborhoodOperator describes the behavior to perform neighborhoods structures.
//
// The Neighborhood N('s) of a candidate solution 's E S is defined by all candidates solutions
// that can be obtained by applying specific modification or movements to 's.
// The Neighborhood structure plays a crucial role in the performance. If the N('s) structure
// is not adequate to the problem, any s-metaheuristic will fail to solve the problem.
//
// A Neighborhood function N is a mapping N: S -> 2's that assigns to each solution 's of S
// a set fo solutions N('s) < S such a ”s E N('s). A solution ”s in the neighborhood of
// 's (”s E N('s)) is called a neighbor of 's. A neighbord is generated by the application of
// a move operator 'm that performs a small perturbation to the solution 's. The structure of
// the N('s) depends on the target optimization problem.
type NeighborhoodOperator interface {
	// BestImprovement chooses the best solution from N('s) where all possible moves are performed
	BestImprovement(s *Solution) *Solution
	// The first improvement strategy tries to avoid the time complexity of exploring
	FirstImprovement(s *Solution) *Solution
}

// Shake transforms one complete candidate solution into another
// complete candidate solution.
type Shaker interface {
	Shake(s *Solution, r *Solution)
}