package vns

// BasicVNS implements a metaheuristic method for solving a set of combinatorial optimization
// and global optimization problems. It explores distant neighborhoods of the current incumbent solution
// and moves from there to a new one if and only if an improvement was made. The local search method is
// applied repeatedly to get from solutions in the neighborhood to local optima.

// The method uses an initial solution (random, greedy), a Local Search neighborhood operator,
// and a set of Shaker's operators that will be used to find new solutions with the Local Search algorithm.
func BasicVNS(s *Solution, operator NeighborhoodOperator, shakers ...Shaker) *Solution {
	currSol := s
	kmax := len(shakers)
	maxIter := 100
	for i := 0; i < maxIter; i++ {
		k := 0
		// For each shaker computes LocalSearch algorithm to improve
		// the current solution.
		for k < kmax {
			shakeSol := &Solution{}
			shakers[k].Shake(currSol, shakeSol)
			// LocalSerch implementation in order to get a local optima solution
			// using a shaker and the neighboorhord operator.
			localOptimumSol := LocalSearch(shakeSol, operator)
			if localOptimumSol.Fitness() > currSol.Fitness() {
				currSol = localOptimumSol
				k = 0
			} else {
				k++
			}
		}
	}
	return currSol
}
