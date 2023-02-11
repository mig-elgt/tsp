package main

import (
	"flag"
	"net"

	"github.com/mig-elgt/tsp/vns/grpc"
	"github.com/mig-elgt/tsp/vns/optimize"
	"github.com/sirupsen/logrus"
)

func main() {
	grpcAddr := flag.String("listen", ":8080", "address:port to listen on")
	flag.Parse()

	lis, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		logrus.Fatalf("could not listen to port %v: %v", *grpcAddr, err)
	}
	logrus.Infof("GRPC listening on %s", *grpcAddr)

	s := grpc.NewAPI(optimize.NewVNS())
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("grpc Server failed: %v", err)
	}
}
