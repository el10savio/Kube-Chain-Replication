package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"../chain"
	"../store"
)

// SetValue ...
func SetValue(w http.ResponseWriter, r *http.Request) {
	var store store.Entry
	_ = json.NewDecoder(r.Body).Decode(&store)

	HEAD, err := chain.GetHead()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to get head node")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseStatus, err := chain.HeadSetValue(HEAD, store)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to set head value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
