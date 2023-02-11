package vns

import (
	"context"
	"reflect"
	"testing"

	"github.com/mig-elgt/tsp/optimizer"
	"github.com/mig-elgt/tsp/optimizer/mocks"
	pb "github.com/mig-elgt/tsp/proto/vns"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_tableGetDistanceMatrix(t *testing.T) {
	type args struct {
		stops          []*optimizer.Stop
		optimizeFnMock func(ctx context.Context, in *pb.OptimizeRequest) (*pb.OptimizeResponse, error)
	}
	testCases := map[string]struct {
		args    args
		wantErr bool
		want    []*optimizer.Stop
	}{
		"service not available": {
			args: args{
				stops: []*optimizer.Stop{},
				optimizeFnMock: func(_ context.Context, _ *pb.OptimizeRequest) (*pb.OptimizeResponse, error) {
					return nil, status.Errorf(codes.Internal, "could not compute table")
				},
			},
			wantErr: true,
		},
		"base case": {
			args: args{
				stops: []*optimizer.Stop{
					{StopID: 1, Location: &optimizer.Location{Lat: 10, Lng: -10}},
					{StopID: 2, Location: &optimizer.Location{Lat: 20, Lng: -20}},
					{StopID: 3, Location: &optimizer.Location{Lat: 30, Lng: -30}},
				},
				optimizeFnMock: func(_ context.Context, _ *pb.OptimizeRequest) (*pb.OptimizeResponse, error) {
					return &pb.OptimizeResponse{
						Stops: []*pb.Stop{
							{ID: 2},
							{ID: 1},
							{ID: 3},
						},
					}, nil
				},
			},
			want: []*optimizer.Stop{
				{StopID: 2, Location: &optimizer.Location{Lat: 20, Lng: -20}},
				{StopID: 1, Location: &optimizer.Location{Lat: 10, Lng: -10}},
				{StopID: 3, Location: &optimizer.Location{Lat: 30, Lng: -30}},
			},
			wantErr: false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			svc := vns{
				client: &mocks.BasicVNSServiceClientMock{
					OptimizeFn: tc.args.optimizeFnMock,
				},
			}
			got, _, err := svc.Optimize(tc.args.stops, []float64{})
			if (err != nil) != tc.wantErr {
				t.Fatalf("got %v; want %v", err, tc.wantErr)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
