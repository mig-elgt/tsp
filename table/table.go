package table

import "net/http"

// Locations represents a destination location
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Cost describes the costs between two locations
type Cost struct {
	// Duration between two locations in seconds
	Duration float64 `json:"duration"`
	// Distance between two locations in meters
	Distance float64 `json:"distance"`
}

// TableService describes the behavior to performs a Distance Matrix object
type TableService interface {
	// Fetch fetches a cost distance matrix for a set of locations
	Fetch(locations []*Location) ([][]*Cost, error)
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
