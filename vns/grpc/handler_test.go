package grpc

import (
	"context"
	"errors"
	"reflect"
	"testing"

	pb "github.com/mig-elgt/tsp/proto/vns"
	"github.com/mig-elgt/tsp/vns"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestHandler_optimizeRoute(t *testing.T) {
	type args struct {
		OptimizeFnMock func(cluster *vns.Cluster) ([]vns.Stop, error)
		request        *pb.OptimizeRequest
	}
	testCases := map[string]struct {
		args           args
		wantStatusCode codes.Code
		wantResponse   *pb.OptimizeResponse
	}{
		"empty stops": {
			args: args{
				OptimizeFnMock: func(_ *vns.Cluster) ([]vns.Stop, error) {
					return []vns.Stop{}, errors.New("got empty list of stops")
				},
				request: &pb.OptimizeRequest{
					Stops: []*pb.Stop{},
				},
			},
			wantStatusCode: codes.InvalidArgument,
		},
		"base case": {
			args: args{
				request: &pb.OptimizeRequest{
					Stops: []*pb.Stop{
						{
							ID: 1,
						},
						{
							ID: 2,
						},
						{
							ID: 3,
						},
					},
					Matrix: []float64{0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
				OptimizeFnMock: func(_ *vns.Cluster) ([]vns.Stop, error) {
					return []vns.Stop{
						{
							StopID:   3,
							Location: &vns.Location{},
						},
						{
							StopID:   2,
							Location: &vns.Location{},
						},
						{
							StopID:   1,
							Location: &vns.Location{},
						},
					}, nil
				},
			},
			wantStatusCode: codes.OK,
			wantResponse: &pb.OptimizeResponse{
				Stops: []*pb.Stop{
					{ID: 3},
					{ID: 2},
					{ID: 1},
				},
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			h := &handler{
				optimizer: &OptimizerMock{
					OptimizeFn: tc.args.OptimizeFnMock,
				},
			}
			got, err := h.optimizeRoute(context.Background(), tc.args.request)
			s, _ := status.FromError(err)
			if got, want := s.Code(), tc.wantStatusCode; got != want {
				t.Fatalf("optimizeRoute() got status code %v; want %v", got, tc.wantStatusCode)
			}
			if !reflect.DeepEqual(got, tc.wantResponse) {
				t.Fatalf("OptimizeRoute() got  %v; want %v", got, tc.wantResponse)
			}
		})
	}
}

type OptimizerMock struct {
	OptimizeFn func(cluster *vns.Cluster) ([]vns.Stop, error)
}

func (op *OptimizerMock) Optimize(cluster *vns.Cluster) ([]vns.Stop, error) {
	return op.OptimizeFn(cluster)
}

func TestHandler_createCluster(t *testing.T) {
	type args struct {
		request *pb.OptimizeRequest
	}
	testCases := map[string]struct {
		args args
		want *vns.Cluster
	}{
		"two stops": {
			args: args{
				request: &pb.OptimizeRequest{
					Stops: []*pb.Stop{
						{
							ID:  1,
							Lat: 100,
							Lng: 100,
						},
						{
							ID:  2,
							Lat: 200,
							Lng: 200,
						},
					},
					Matrix: []float64{0, 10, 20, 0},
				},
			},
			want: &vns.Cluster{
				Stops: []vns.Stop{
					{
						StopID: 1,
						Location: &vns.Location{
							Lat: 100,
							Lng: 100,
						},
					},
					{
						StopID: 2,
						Location: &vns.Location{
							Lat: 200,
							Lng: 200,
						},
					},
				},
				CostMatrix: vns.CostMatrix{
					{
						{Distance: 0},
						{Distance: 10},
					},
					{
						{Distance: 20},
						{Distance: 0},
					},
				},
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			h := &handler{}
			if got := h.createCluster(tc.args.request); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("createCluster(req) got %v; want %v", got, tc.want)
			}
		})
	}
}
