package optimizer

// Location describes a location point.
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Stop describes a customer's stop location.
type Stop struct {
	Name     string    `json:"name"`
	StopID   int       `json:"stop_id,omitempty"`
	Location *Location `json:"location"`
}

// TableService describes the behavior to get the Distance Matrix for set of stops
type TableService interface {
	GetDistanceMatrix(stops []*Stop) ([]float64, error)
}

// VNSService describes an inteface to perform a route optimization algorithm
type VNSService interface {
	Optimize(stops []*Stop, matrix []float64) ([]*Stop, float64, error)
}
