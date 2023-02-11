package neighborhood

import "github.com/mig-elgt/tsp/vns"

type SolutionMover interface {
	Evaluate(s *vns.Solution, k int) float64
}

func EvaluateMovesAndReturnTheBestOne(moves []SolutionMover, s *vns.Solution, k int) SolutionMover {
	bestMove := moves[0]
	bestCost := bestMove.Evaluate(s, k)
	for i := 1; i < len(moves); i++ {
		newCost := moves[i].Evaluate(s, k)
		if newCost < bestCost {
			bestCost = newCost
			bestMove = moves[i]
		}
	}
	return bestMove
}
