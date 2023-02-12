package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/mig-elgt/tsp/optimizer"
	"github.com/mig-elgt/tsp/optimizer/mocks"
	"github.com/pkg/errors"
)

func Test_handlerTSP(t *testing.T) {
	type args struct {
		body                    *bytes.Buffer
		getDistanceMatrixFnMock func(stops []*optimizer.Stop) ([]float64, error)
		optimizeFnMock          func(stops []*optimizer.Stop, matrix []float64) ([]*optimizer.Stop, float64, error)
	}
	testCases := map[string]struct {
		args           args
		wantStatusCode int
		wantResponse   []byte
	}{
		"body request bad format": {
			args: args{
				body: bytes.NewBufferString(`this is a plain text`),
			},
			wantStatusCode: http.StatusBadRequest,
			wantResponse:   []byte("{\"error\":{\"status\":400,\"error\":\"INVALID_ARGUMENT\",\"description\":\"The request body entity is bad format.\"}}\n"),
		},
		"table service is not available": {
			args: args{
				body: bytes.NewBufferString(`{"stops":[{"name":"foo","location":{"lat":100,"lng":-100}}]}`),
				getDistanceMatrixFnMock: func(_ []*optimizer.Stop) ([]float64, error) {
					return nil, errors.New("table service is shut down")
				},
			},
			wantStatusCode: http.StatusInternalServerError,
			wantResponse:   []byte("{\"error\":{\"status\":500,\"error\":\"INTERNAL\",\"description\":\"Something went wrong...\"}}\n"),
		},
		"vns service is not available": {
			args: args{
				body: bytes.NewBufferString(`{"stops":[{"name":"foo","location":{"lat":100,"lng":-100}}]}`),
				getDistanceMatrixFnMock: func(_ []*optimizer.Stop) ([]float64, error) {
					return []float64{0, 1, 3, 0}, nil
				},
				optimizeFnMock: func(stops []*optimizer.Stop, matrix []float64) ([]*optimizer.Stop, float64, error) {
					return nil, 0, errors.New("vns service is shutd down")
				},
			},
			wantStatusCode: http.StatusInternalServerError,
			wantResponse:   []byte("{\"error\":{\"status\":500,\"error\":\"INTERNAL\",\"description\":\"Something went wrong...\"}}\n"),
		},
		"base case": {
			args: args{
				body: bytes.NewBufferString(`{"stops":[{"name":"foo","location":{"lat":100,"lng":-100}},{"name":"bar","location":{"lat":200,"lng":-200}}]}`),
				getDistanceMatrixFnMock: func(_ []*optimizer.Stop) ([]float64, error) {
					return []float64{0, 1, 3, 0}, nil
				},
				optimizeFnMock: func(stops []*optimizer.Stop, matrix []float64) ([]*optimizer.Stop, float64, error) {
					return []*optimizer.Stop{
						{
							Name:   "bar",
							StopID: 2,
							Location: &optimizer.Location{
								Lat: 200,
								Lng: -200,
							},
						},
						{
							Name:   "foo",
							StopID: 1,
							Location: &optimizer.Location{
								Lat: 100,
								Lng: -100,
							},
						},
					}, 100, nil
				},
			},
			wantStatusCode: http.StatusOK,
			wantResponse:   []byte("{\"route\":[{\"name\":\"bar\",\"stop_id\":2,\"location\":{\"lat\":200,\"lng\":-200}},{\"name\":\"foo\",\"stop_id\":1,\"location\":{\"lat\":100,\"lng\":-100}}],\"total_distance\":100}\n"),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/v1/tsp", tc.args.body)
			if err != nil {
				t.Fatal(err)
			}
			rec := httptest.NewRecorder()
			h := handler{
				table: &mocks.TableServiceMock{
					GetDistanceMatrixFn: tc.args.getDistanceMatrixFnMock,
				},
				vns: &mocks.VNSServiceMock{
					OptimizeFn: tc.args.optimizeFnMock,
				},
			}
			h.TSP(rec, req)
			if got, want := rec.Code, tc.wantStatusCode; got != want {
				t.Fatalf("%v: TSP(w,r) status code got %v; want %v", name, got, want)
			}
			if rec.Result().StatusCode == http.StatusOK {
				if got, want := rec.Body.Bytes(), tc.wantResponse; !reflect.DeepEqual(got, want) {
					t.Errorf("%v: TSP(w, r) body response got \n%v; want \n%v", name, string(got), string(want))
				}
			}
		})
	}

}
