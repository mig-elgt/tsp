package mocks

import (
	"context"

	"github.com/mig-elgt/tsp/optimizer"
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

type TableServiceMock struct {
	GetDistanceMatrixFn func(stops []*optimizer.Stop) ([]float64, error)
}

func (t *TableServiceMock) GetDistanceMatrix(stops []*optimizer.Stop) ([]float64, error) {
	return t.GetDistanceMatrixFn(stops)
}

type VNSServiceMock struct {
	OptimizeFn func(stops []*optimizer.Stop, matrix []float64) ([]*optimizer.Stop, float64, error)
}

func (v *VNSServiceMock) Optimize(stops []*optimizer.Stop, matrix []float64) ([]*optimizer.Stop, float64, error) {
	return v.OptimizeFn(stops, matrix)
}
