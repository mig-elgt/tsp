package mocks

import (
	"net/http"

	"github.com/mig-elgt/tsp/table"
)

type HTTPClientMock struct {
	DoFn func(req *http.Request) (*http.Response, error)
}

func (c *HTTPClientMock) Do(req *http.Request) (*http.Response, error) {
	return c.DoFn(req)
}

type TableServiceMock struct {
	FetchFn func(locations []*table.Location) ([][]*table.Cost, error)
}

// Fetch fetches a cost distance matrix for a set of locations
func (t *TableServiceMock) Fetch(locations []*table.Location) ([][]*table.Cost, error) {
	return t.FetchFn(locations)
}
