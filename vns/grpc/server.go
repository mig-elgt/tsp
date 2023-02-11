package grpc

import (
	"context"

	pb "github.com/mig-elgt/tsp/proto/vns"
	"github.com/mig-elgt/tsp/vns"
	"google.golang.org/grpc"
)

type grpcServer struct {
	pb.UnimplementedBasicVNSServiceServer
	handler
}

func NewApi(optimizer vns.Optimizer) *grpc.Server {
	rootServer := grpc.NewServer()
	s := &grpcServer{
		handler: handler{optimizer},
	}
	pb.RegisterBasicVNSServiceServer(rootServer, s)
	return rootServer
}

func (s *grpcServer) Optimize(ctx context.Context, req *pb.OptimizeRequest) (*pb.OptimizeResponse, error) {
	return s.handler.optimizeRoute(ctx, req)
}
