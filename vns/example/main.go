package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mig-elgt/tsp/vns"
	"google.golang.org/grpc"

	pb "github.com/mig-elgt/tsp/proto/vns"
)

func main() {
	cmFileName := flag.String("file", "file", "file name of cost matrix json")
	address := flag.String("h", ":8080", "address for the service")
	flag.Parse()

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not get connection to address %s: %v", *address, err)
	}
	client := pb.NewBasicVNSServiceClient(conn)

	cm, err := getCostMatrix(*cmFileName)
	if err != nil {
		log.Fatalf("could not get cost matrix %v: %v", *cmFileName, err)
	}

	var matrix []float64
	for i := 0; i < len(cm); i++ {
		for j := 0; j < len(cm[0]); j++ {
			matrix = append(matrix, cm[i][j].Distance)
		}
	}
	stops := getStops(len(cm) - 1)
	optimizeRoute(client, stops, matrix)
}

func optimizeRoute(client pb.BasicVNSServiceClient, stops []*pb.Stop, matrix []float64) {
	resp, err := client.Optimize(context.Background(), &pb.OptimizeRequest{
		Stops:  stops,
		Matrix: matrix,
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, stop := range resp.Stops {
		fmt.Printf("%v ", stop.ID)
	}
	fmt.Println(resp.TotalDistance)
}

type costMatrixJSON struct {
	CostMatrix vns.CostMatrix `json:"matrix"`
}

func getCostMatrix(fileName string) (vns.CostMatrix, error) {
	cm := costMatrixJSON{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&cm); err != nil {
		return nil, err
	}
	return cm.CostMatrix, nil
}

func getStops(stopsSize int) []*pb.Stop {
	stops := []*pb.Stop{}
	for i := 0; i < stopsSize; i++ {
		stops = append(stops,
			&pb.Stop{
				ID: int32(i + 1),
			})
	}
	return stops
}
