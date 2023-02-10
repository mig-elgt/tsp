package vns

// Location describes a location point.
type Location struct {
	Lat float64
	Lng float64
}

// Stop describes a customer's stop location.
type Stop struct {
	StopID   int
	Location *Location
}

// TravelCost holds the main properties to describes a cost between two points.
type TravelCost struct {
	Distance float32 `json:"distance_in_meters"`
}

// CostMatrix describes a distance matrix between all stops.
type CostMatrix [][]TravelCost

// Clusters defines an abstraction data structure to store all data
// in order to compute a route optimization.
type Cluster struct {
	Stops      []Stop
	CostMatrix CostMatrix
}

// Optimizer describes an inteface to compute an algorithm
// to solve the VRP problem.
type Optimizer interface {
	Optimize(cluster *Cluster) ([]Stop, error)
}
