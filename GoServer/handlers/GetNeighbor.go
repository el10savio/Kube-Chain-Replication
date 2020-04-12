package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// GetNeighbor is the handler interface to
// obtain the corresponding neighbor
// for the requested node
func GetNeighbor(w http.ResponseWriter, r *http.Request) {
	neighbor := getNeighbor()

	// Check if neighbor is empty
	if neighbor == "" {
		log.Error("neighbor not present")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log.WithFields(log.Fields{"Neighbor": neighbor}).Debug("successfull GetNeighbor")
	json.NewEncoder(w).Encode(neighbor)
}
