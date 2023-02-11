package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "github.com/mig-elgt/tsp/proto/table"
)

func main() {
	address := flag.String("h", ":8080", "address for the service")
	flag.Parse()

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not get connection to address %s: %v", *address, err)
	}

	stops := []*pb.Stop{
		{Lat: 52.517033, Lng: 13.388798},
		{Lat: 52.529432, Lng: 13.39763},
		{Lat: 52.523239, Lng: 13.428554},
	}
	client := pb.NewTableServiceClient(conn)

	fetchDistanceMatrix(client, stops)
}

func fetchDistanceMatrix(client pb.TableServiceClient, stops []*pb.Stop) {
	resp, err := client.Fetch(context.Background(), &pb.FetchRequest{
		Stops: stops,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Matrix)
}
