package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Neighbor ...
var Neighbor string

// UpdateNeighbor ...
func UpdateNeighbor(w http.ResponseWriter, r *http.Request) {
	neighbor := mux.Vars(r)["neighbor"]

	updateNeighbor(neighbor)

	if neighbor != Neighbor {
		log.WithFields(log.Fields{
			"local neighbor":    neighbor,
			"received neighbor": Neighbor,
		}).Error("failed updating neighbor")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"Neighbor": Neighbor,
	}).Debug("successfull UpdateNeighbor")

	w.WriteHeader(http.StatusOK)
}

// updateNeighbor ...
func updateNeighbor(neighbor string) {
	Neighbor = neighbor
}

// getNeighbor ...
func getNeighbor() string {
	return Neighbor
}
