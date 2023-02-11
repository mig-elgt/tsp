package osrm

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kr/pretty"
	"github.com/mig-elgt/tsp/table"
	"github.com/pkg/errors"
)

const osrmURLBase = "http://router.project-osrm.org"

type osrm struct {
	client table.HTTPClient
}

func New() *osrm {
	return &osrm{
		client: &http.Client{},
	}
}

// Fetch fetches a cost distance matrix for a set of locations
func (o *osrm) Fetch(locations []*table.Location) ([][]*table.Cost, error) {
	// Prepare request
	url := fmt.Sprintf("%v/table/v1/driving/%v?annotations=distance,duration", osrmURLBase, o.convertLocationsToStringFormat(locations))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not create request")
	}
	// Send requet to OSRM Server
	resp, err := o.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "could not send request")
	}
	defer resp.Body.Close()
	// Validate status code and handle response errors
	if resp.StatusCode != http.StatusOK {
		var respErr struct {
			Message string `json:"message"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&respErr); err != nil {
			return nil, errors.Wrap(err, "could not decode error body response")
		}
		return nil, fmt.Errorf("web server error: %s", respErr.Message)
	}
	// Decode response body
	var result struct {
		Durations [][]float64 `json:"durations"`
		Distances [][]float64 `json:"distances"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, errors.Wrap(err, "could not decode result body response")
	}
	pretty.Print(result)
	// Convert result response to distance matrix object
	matrix := [][]*table.Cost{}
	for i := 0; i < len(locations); i++ {
		cm := []*table.Cost{}
		for j := 0; j < len(locations); j++ {
			cm = append(cm, &table.Cost{
				Distance: result.Distances[i][j],
				Duration: result.Durations[i][j],
			})
		}
		matrix = append(matrix, cm)
	}
	return matrix, nil
}

func (o *osrm) convertLocationsToStringFormat(stops []*table.Location) string {
	locString := ""
	size := len(stops)
	for i := 0; i < size-1; i++ {
		locString += fmt.Sprintf("%v,%v;", stops[i].Lng, stops[i].Lat)
	}
	return fmt.Sprintf("%v%v,%v", locString, stops[size-1].Lng, stops[size-1].Lat)
}
