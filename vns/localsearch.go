package vns

// LocalSearch algorithm moves from solution to solution in the space of candidate solutions (the search space) by applying local changes using a neighborhood operator. The algorithm finishes when it finds a local optima solution.
// A local optima solution is found when the fitness value of the new solution is not better than the current solution.
func LocalSearch(s *Solution, opt NeighborhoodOperator) *Solution {
	currSol := s
	for {
		newSol := opt.BestImprovement(currSol)
		if newSol.Fitness() <= currSol.Fitness() {
			break
		}
		currSol = newSol
	}
	return currSol
}
