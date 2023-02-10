package neighborhood

import (
	"testing"

	"github.com/mig-elgt/tsp/vns"
)

func Test_EvaluteMovesAndReturnTheBestOne(t *testing.T) {
	testCases := map[string]struct {
		moves []SolutionMover
		want  swapMoveMock
	}{
		"base case": {
			moves: []SolutionMover{
				swapMoveMock{
					cost: 10,
					EvaluateFn: func(_ *vns.Solution, _ int) float32 {
						return 10
					},
				},
				swapMoveMock{
					cost: 20,
					EvaluateFn: func(_ *vns.Solution, _ int) float32 {
						return 20
					},
				},
				swapMoveMock{
					cost: 70,
					EvaluateFn: func(_ *vns.Solution, _ int) float32 {
						return 70
					},
				},
			},
			want: swapMoveMock{
				cost: 10,
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := EvaluateMovesAndReturnTheBestOne(tc.moves, nil, 1)
			m := got.(swapMoveMock)
			if m.cost != tc.want.cost {
				t.Fatalf("got %v, want %v", m, tc.want)
			}
		})
	}
}

type swapMoveMock struct {
	cost       float32
	EvaluateFn func(s *vns.Solution, k int) float32
}

func (sw swapMoveMock) Evaluate(s *vns.Solution, k int) float32 {
	return sw.EvaluateFn(s, k)
}
