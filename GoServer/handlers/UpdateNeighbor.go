package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Neighbor is the IP address
// of the neighbor node
var Neighbor string

// UpdateNeighbor is the handler interface to
// update the neighbor value from GoProxy
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

// updateNeighbor updates the  local 
// package neighbor value
func updateNeighbor(neighbor string) {
	Neighbor = neighbor
}

// getNeighbor returns the  local 
// package neighbor value
func getNeighbor() string {
	return Neighbor
}
