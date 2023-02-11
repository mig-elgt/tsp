package grpc

import (
	"context"

	pb "github.com/mig-elgt/tsp/proto/table"
	"github.com/mig-elgt/tsp/table"
	"google.golang.org/grpc"
)

type grpcServer struct {
	pb.UnimplementedTableServiceServer
	handler
}

// NewAPI create new grpc Server and register the VNS service
// with the optimizer handle.
func NewAPI(t table.TableService) *grpc.Server {
	rootServer := grpc.NewServer()
	s := &grpcServer{
		handler: handler{svc: t},
	}
	pb.RegisterTableServiceServer(rootServer, s)
	return rootServer
}

// Optimize optimizes a set of stops and returns an optima route.
func (s *grpcServer) Fetch(ctx context.Context, req *pb.FetchRequest) (*pb.FetchResponse, error) {
	return s.handler.fetchDistanceMatrix(ctx, req)
}
