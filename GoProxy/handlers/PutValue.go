package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/el10savio/Kube-Chain-Replication/GoProxy/chain"
	"github.com/el10savio/Kube-Chain-Replication/GoProxy/store"
)

// SetValue handler is a proxy
// to send a write request to
// the HEAD node of the chain
func SetValue(w http.ResponseWriter, r *http.Request) {
	// json decode the KV pair
	var store store.Entry
	_ = json.NewDecoder(r.Body).Decode(&store)

	// Get the HEAD node of the chain
	HEAD, err := chain.GetHead()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to get head node")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send a write request to the HEAD
	responseStatus, err := chain.HeadSetValue(HEAD, store)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to set head value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check if a valid response is
	// sent back from the HEAD
	if responseStatus != http.StatusOK {
		log.WithFields(log.Fields{"status": responseStatus}).Error("invalid response from head")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"key":   store.Key,
		"value": store.Value,
	}).Debug("successfull SetValue")

	w.WriteHeader(http.StatusOK)
}
