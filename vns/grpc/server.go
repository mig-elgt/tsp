package grpc

import (
	"context"

	pb "github.com/mig-elgt/tsp/proto/vns"
	"google.golang.org/grpc"
)

type grpcServer struct {
	pb.UnimplementedBasicVNSServiceServer
	handler
}

func NewApi() *grpc.Server {
	rootServer := grpc.NewServer()
	s := &grpcServer{
		handler: handler{},
	}
	pb.RegisterBasicVNSServiceServer(rootServer, s)
	return rootServer
}

func (g *grpcServer) Optimize(_ context.Context, _ *pb.OptimizeRequest) (*pb.OptimizeResponse, error) {
	panic("not implemented") // TODO: Implement
}
