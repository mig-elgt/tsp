package table

import (
	"context"
	"reflect"
	"testing"

	"github.com/mig-elgt/tsp/optimizer"
	"github.com/mig-elgt/tsp/optimizer/mocks"
	pb "github.com/mig-elgt/tsp/proto/table"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_tableGetDistanceMatrix(t *testing.T) {
	type args struct {
		stops       []*optimizer.Stop
		fetchFnMock func(ctx context.Context, in *pb.FetchRequest) (*pb.FetchResponse, error)
	}
	testCases := map[string]struct {
		args    args
		wantErr bool
		want    []float64
	}{
		"service not available": {
			args: args{
				stops: []*optimizer.Stop{},
				fetchFnMock: func(_ context.Context, _ *pb.FetchRequest) (*pb.FetchResponse, error) {
					return nil, status.Errorf(codes.Internal, "could not compute table")
				},
			},
			wantErr: true,
		},
		"base case": {
			args: args{
				stops: []*optimizer.Stop{
					{
						Location: &optimizer.Location{},
					},
					{
						Location: &optimizer.Location{},
					},
				},
				fetchFnMock: func(_ context.Context, _ *pb.FetchRequest) (*pb.FetchResponse, error) {
					return &pb.FetchResponse{
						Matrix: []float64{0, 1, 2, 1, 0, 2, 1, 2, 0},
					}, nil
				},
			},
			wantErr: false,
			want:    []float64{0, 1, 2, 1, 0, 2, 1, 2, 0},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			svc := table{
				client: &mocks.TableServiceClientMock{
					FetchFn: tc.args.fetchFnMock,
				},
			}
			got, err := svc.GetDistanceMatrix(tc.args.stops)
			if (err != nil) != tc.wantErr {
				t.Fatalf("got %v; want %v", err, tc.wantErr)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
