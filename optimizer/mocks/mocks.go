package mocks

import (
	"context"

	"github.com/mig-elgt/tsp/proto/table"
	pb "github.com/mig-elgt/tsp/proto/table"
	"google.golang.org/grpc"
)

type TableServiceClientMock struct {
	FetchFn func(ctx context.Context, in *pb.FetchRequest) (*pb.FetchResponse, error)
}

func (t *TableServiceClientMock) Fetch(ctx context.Context, in *table.FetchRequest, opts ...grpc.CallOption) (*table.FetchResponse, error) {
	return t.FetchFn(ctx, in)
}
