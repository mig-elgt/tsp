package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mig-elgt/sender"
	"github.com/mig-elgt/sender/codes"
	"github.com/mig-elgt/tsp/optimizer"
	"github.com/sirupsen/logrus"
)

type handler struct {
	table optimizer.TableService
	vns   optimizer.VNSService
}

// New creates a new HTTP Handler
func New(table optimizer.TableService, vns optimizer.VNSService) http.Handler {
	r := mux.NewRouter()
	h := handler{table, vns}
	r.HandleFunc("/api/v1/tsp", h.TSP).Methods("POST")
	return r
}

// TSP represents a handle for the request POST /api/v1/tsp
func (h handler) TSP(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Stops []*optimizer.Stop `json:"stops"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logrus.Errorf("could not decode request body: %v", err)
		sender.
			NewJSON(w, http.StatusBadRequest).
			WithError(codes.InvalidArgument, "Request bad format object").Send()
		return
	}

	matrix, err := h.table.GetDistanceMatrix(request.Stops)
	if err != nil {
		logrus.Errorf("could not get distance matrix: %v", err)
		sender.
			NewJSON(w, http.StatusInternalServerError).
			WithError(codes.Internal, "Something went wrong...").Send()
		return
	}

	stops, distance, err := h.vns.Optimize(request.Stops, matrix)
	if err != nil {
		logrus.Errorf("could not optimize route: %v", err)
		sender.
			NewJSON(w, http.StatusInternalServerError).
			WithError(codes.Internal, "Something went wrong...").Send()
		return
	}

	type tspResponse struct {
		Route         []*optimizer.Stop `json:"route"`
		TotalDistance float64           `json:"total_distance"`
	}
	sender.
		NewJSON(w, http.StatusOK).
		Send(&tspResponse{
			Route:         stops,
			TotalDistance: distance,
		})
}
