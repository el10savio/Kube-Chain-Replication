package handlers

import (
	"encoding/json"
	"net/http"

	"../chain"

	log "github.com/sirupsen/logrus"
)

// GetPods handler returns the active goredis pods 
func GetPods(w http.ResponseWriter, r *http.Request) {
	// Get active goredis pods 
	pods, err := chain.GetPodsList()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed getting pods from cluster")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"pods": pods,
	}).Debug("successfull GetPods")

	// json encoded pods list
	json.NewEncoder(w).Encode(pods)
}
