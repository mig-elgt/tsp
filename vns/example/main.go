package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mig-elgt/tsp/vns"
	"github.com/mig-elgt/tsp/vns/optimize"
)

func main() {
	cmFileName := flag.String("file", "file", "file name of cost matrix json")
	fleeSize := flag.Int("vehicles", 1, "fleet size")
	flag.Parse()
	cm, err := getCostMatrix(*cmFileName)
	if err != nil {
		log.Fatalf("could not get cost matrix %v: %v", *cmFileName, err)
	}
	stops := getStops(len(cm) - 1)
	vehicles := []vns.Vehicle{}
	for i := 0; i < *fleeSize; i++ {
		vehicles = append(vehicles, vns.Vehicle{
			Capacity: 50, StartLocation: &vns.Location{Name: "Foo Bar", Lat: 22.153458, Lng: -100.977310},
		})
	}
	cluster := &vns.Cluster{
		CostMatrix: cm,
		Stops:      stops,
		Vehicles:   vehicles,
	}
	vns := optimize.NewVNS()
	result, err := vns.Optimize(cluster)
	if err != nil {
		log.Fatalf("could not optimize route: %v", err)
	}
	fmt.Println(result)
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

func getStops(stopsSize int) []vns.Stop {
	stops := []vns.Stop{}
	for i := 0; i < stopsSize; i++ {
		stops = append(stops, vns.Stop{
			Demand: 1,
			StopID: i + 1,
		})
	}
	return stops
}
