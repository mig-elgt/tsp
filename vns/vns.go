package vns

// Location describes a location point.
type Location struct {
	Name string  `json:"company_name" valid:"required~The name is required"`
	Lat  float64 `json:"lat"`
	Lng  float64 `json:"lng"`
}

// Stop describes a customer's stop location.
type Stop struct {
	StopID              int
	StopName            string
	Demand              float32
	Location            *Location
	StartTimeWindow     float32
	EndTimeWindow       float32
	DurationServiceTime int64 // in minutes
	ArrivalTime         float32
	WaitingTime         float32
	DepartureTime       float32
}

// TravelCost holds the main properties to describes a cost between two points.
type TravelCost struct {
	Distance float32 `json:"distance_in_meters"`
	Time     float32 `json:"travel_time_in_minutes"`
}

// CostMatrix describes a distance matrix between all stops.
type CostMatrix [][]TravelCost

// Vehicle describes a vehicle data.
type Vehicle struct {
	VehicleID        string
	Capacity         float32   `json:"capacity"`
	StartLocation    *Location `json:"start_location"`
	ShiftStart       string    `json:"shift_start"`
	ShiftStartNumber float32   `json:"shift_start_number"`
	ShiftEnd         string    `json:"shift_end"`
	ShiftEndNumber   float32   `json:"shift_end_number"`
}

// Clusters defines an abstraction data structure to store all data
// in order to compute a route optimization.
type Cluster struct {
	Depot      Stop
	Stops      []Stop
	Vehicles   []Vehicle
	CostMatrix CostMatrix
}
