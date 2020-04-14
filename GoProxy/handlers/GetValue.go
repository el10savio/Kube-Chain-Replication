package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/el10savio/Kube-Chain-Replication/GoProxy/chain"
)

// GetValue handler is a proxy 
// to send a read request to 
// the TAIL node of the chain
func GetValue(w http.ResponseWriter, r *http.Request) {
	// Obtain the key from URL params
	key := mux.Vars(r)["key"]

	// Get the TAIL node of the chain
	TAIL, err := chain.GetTail()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to get tail node")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send a read request to the TAIL
	store, err := chain.TailGetValue(TAIL, key)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to Get tail value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"key":   store.Key,
		"value": store.Value,
	}).Debug("successfull GetValue")

	// json encode response KV pair
	json.NewEncoder(w).Encode(store)
}
