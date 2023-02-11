package optimizer

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

// TableService describes the behavior to get the Distance Matrix for set of stops
type TableService interface {
	GetDistanceMatrix(stops []*Stop) ([]float64, error)
}

// VNSService describes an inteface to perform a route optimization algorithm
type VNSService interface {
	Optimize(stops []*Stop, matrix []float64) ([]*Stop, float64, error)
}
