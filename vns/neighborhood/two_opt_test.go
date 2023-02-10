package neighborhood

import (
	"reflect"
	"testing"

	"github.com/mig-elgt/tsp/vns"
)

func TestTwoOpt_generateMoves(t *testing.T) {
	testCases := map[string]struct {
		intRandFnMock func(min, max int) int
		in            *vns.Solution
		want          []SolutionMover
	}{
		"1 route, 4 customers: select route 1": {
			intRandFnMock: func(_, _ int) int {
				return 0
			},
			in: &vns.Solution{
				Customers: []int{1, 2, 3, 4},
				Routes: []vns.Route{
					{Start: 0, Size: 4, Cost: 0},
				},
			},
			want: []SolutionMover{
				&twoOptMove{0, 0, 1},
				&twoOptMove{0, 0, 2},
				&twoOptMove{0, 0, 3},
				&twoOptMove{0, 1, 2},
				&twoOptMove{0, 1, 3},
				&twoOptMove{0, 2, 3},
			},
		},
		"2 routes, 8 customers: select route 2": {
			intRandFnMock: func(_, _ int) int {
				return 1
			},
			in: &vns.Solution{
				Customers: []int{1, 2, 3, 4, 5, 6, 7, 8},
				Routes: []vns.Route{
					{Start: 0, Size: 4, Cost: 0},
					{Start: 4, Size: 4, Cost: 0},
				},
			},
			want: []SolutionMover{
				&twoOptMove{1, 4, 5},
				&twoOptMove{1, 4, 6},
				&twoOptMove{1, 4, 7},
				&twoOptMove{1, 5, 6},
				&twoOptMove{1, 5, 7},
				&twoOptMove{1, 6, 7},
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			to := twoOpt{
				rand: &intRandomMock{
					IntFn: tc.intRandFnMock,
				},
			}
			if got := to.generateMoves(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}

type intRandomMock struct {
	IntFn func(min, max int) int
}

func (i intRandomMock) Int(min, max int) int {
	return i.IntFn(min, max)
}
