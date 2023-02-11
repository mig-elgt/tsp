package grpc

import (
	"context"
	"errors"
	"reflect"
	"testing"

	pb "github.com/mig-elgt/tsp/proto/table"
	"github.com/mig-elgt/tsp/table"
	"github.com/mig-elgt/tsp/table/mocks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestHandler_optimizeRoute(t *testing.T) {
	type args struct {
		FetchFnMock func(locations []*table.Location) ([][]*table.Cost, error)
		request     *pb.FetchRequest
	}
	testCases := map[string]struct {
		args           args
		wantStatusCode codes.Code
		wantResponse   *pb.FetchResponse
	}{
		"empty stops": {
			args: args{
				request: &pb.FetchRequest{
					Stops: []*pb.Stop{},
				},
			},
			wantStatusCode: codes.InvalidArgument,
		},
		"osrm service not available": {
			args: args{
				request: &pb.FetchRequest{
					Stops: []*pb.Stop{
						{},
					},
				},
				FetchFnMock: func(_ []*table.Location) ([][]*table.Cost, error) {
					return nil, errors.New("service not available")
				},
			},
			wantStatusCode: codes.Internal,
		},
		"base case": {
			args: args{
				request: &pb.FetchRequest{
					Stops: []*pb.Stop{
						{},
					},
				},
				FetchFnMock: func(_ []*table.Location) ([][]*table.Cost, error) {
					return [][]*table.Cost{
						{
							&table.Cost{Distance: 10, Duration: 20},
							&table.Cost{Distance: 20, Duration: 30},
						},
						{
							&table.Cost{Distance: 30, Duration: 40},
							&table.Cost{Distance: 50, Duration: 60},
						},
					}, nil
				},
			},
			wantStatusCode: codes.OK,
			wantResponse: &pb.FetchResponse{
				Matrix: []float64{10, 20, 30, 50},
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			h := &handler{
				table: &mocks.TableServiceMock{
					FetchFn: tc.args.FetchFnMock,
				},
			}
			got, err := h.fetchDistanceMatrix(context.Background(), tc.args.request)
			s, _ := status.FromError(err)
			if got, want := s.Code(), tc.wantStatusCode; got != want {
				t.Fatalf("fetchDistanceMatrix(req) got status code %v; want %v", got, tc.wantStatusCode)
			}
			if !reflect.DeepEqual(got, tc.wantResponse) {
				t.Fatalf("fetchDistanceMatrix(req) got  %v; want %v", got, tc.wantResponse)
			}
		})
	}
}
