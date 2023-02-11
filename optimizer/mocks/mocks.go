package mocks

import (
	"context"

	pbt "github.com/mig-elgt/tsp/proto/table"
	pbv "github.com/mig-elgt/tsp/proto/vns"
	"google.golang.org/grpc"
)

type TableServiceClientMock struct {
	FetchFn func(ctx context.Context, in *pbt.FetchRequest) (*pbt.FetchResponse, error)
}

func (t *TableServiceClientMock) Fetch(ctx context.Context, in *pbt.FetchRequest, opts ...grpc.CallOption) (*pbt.FetchResponse, error) {
	return t.FetchFn(ctx, in)
}

type BasicVNSServiceClientMock struct {
	OptimizeFn func(ctx context.Context, in *pbv.OptimizeRequest) (*pbv.OptimizeResponse, error)
}

func (b *BasicVNSServiceClientMock) Optimize(ctx context.Context, in *pbv.OptimizeRequest, opts ...grpc.CallOption) (*pbv.OptimizeResponse, error) {
	return b.OptimizeFn(ctx, in)
}
