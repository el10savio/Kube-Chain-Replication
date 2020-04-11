package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"../chain"
)

// GetValue ...
func GetValue(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]

	TAIL, err := chain.GetTail()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to get tail node")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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

	json.NewEncoder(w).Encode(store)
}
