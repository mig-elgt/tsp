package grpc

import (
	"context"

	pb "github.com/mig-elgt/tsp/proto/table"
	"github.com/mig-elgt/tsp/vns"
	"google.golang.org/grpc"
)

type grpcServer struct {
	pb.UnimplementedTableServiceServer
}

// NewAPI create new grpc Server and register the VNS service
// with the optimizer handle.
func NewAPI(optimizer vns.Optimizer) *grpc.Server {
	rootServer := grpc.NewServer()
	s := &grpcServer{}
	pb.RegisterTableServiceServer(rootServer, s)
	return rootServer
}

// Optimize optimizes a set of stops and returns an optima route.
func (s *grpcServer) Optimize(ctx context.Context, req *pb.FetchRequest) (*pb.FetchResponse, error) {
	panic("not impl")
}
