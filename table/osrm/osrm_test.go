package osrm

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/mig-elgt/tsp/table"
	"github.com/mig-elgt/tsp/table/mocks"
)

func Test_osrmFetch(t *testing.T) {
	testCases := map[string]struct {
		locations []*table.Location
		DoFnMock  func(req *http.Request) (*http.Response, error)
		wantErr   bool
		expectErr error
		want      [][]*table.Cost
	}{
		"could not send request": {
			locations: []*table.Location{
				{},
			},
			DoFnMock: func(_ *http.Request) (*http.Response, error) {
				return nil, errors.New("networking error")
			},
			wantErr:   true,
			expectErr: errors.New("could not send request: networking error"),
		},
		"not route found": {
			locations: []*table.Location{
				{},
			},
			DoFnMock: func(_ *http.Request) (*http.Response, error) {
				json := `{"message": "No Route Found"}`
				r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
				return &http.Response{
					StatusCode: 404,
					Body:       r,
				}, nil
			},
			wantErr:   true,
			expectErr: errors.New("web server error: No Route Found"),
		},
		"base case": {
			locations: []*table.Location{
				{},
				{},
				{},
			},
			DoFnMock: func(_ *http.Request) (*http.Response, error) {
				json := `{"durations":[[1,2,3],[4,5,6],[7,8,9]],"distances":[[10,20,30],[40,50,60],[70,80,90]]}`
				r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
				return &http.Response{
					StatusCode: 200,
					Body:       r,
				}, nil
			},
			wantErr: false,
			want: [][]*table.Cost{
				{
					{Duration: 1, Distance: 10},
					{Duration: 2, Distance: 20},
					{Duration: 3, Distance: 30},
				},
				{
					{Duration: 4, Distance: 40},
					{Duration: 5, Distance: 50},
					{Duration: 6, Distance: 60},
				},
				{
					{Duration: 7, Distance: 70},
					{Duration: 8, Distance: 80},
					{Duration: 9, Distance: 90},
				},
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			svc := osrm{
				client: &mocks.HTTPClientMock{
					DoFn: tc.DoFnMock,
				},
			}
			got, err := svc.Fetch(tc.locations)
			if (err != nil) != tc.wantErr {
				t.Fatalf("case: %v; got %v; want %v", name, err, tc.wantErr)
			}
			if tc.wantErr && !reflect.DeepEqual(err.Error(), tc.expectErr.Error()) {
				t.Fatalf("case: %v; got %v; want %v", name, err, tc.expectErr)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("case: %v; got %v; want %v", name, got, tc.want)
			}
		})
	}

}
