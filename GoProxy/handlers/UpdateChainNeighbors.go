package handlers

import (
	"net/http"

	"../chain"

	log "github.com/sirupsen/logrus"
)

// UpdateChainNeighbors ...
func UpdateChainNeighbors(w http.ResponseWriter, r *http.Request) {
	pods, err := chain.GetPodsList()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed getting pods from cluster")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for index := 0; index < len(pods)-1; index++ {
		pod, neighbor := pods[index], pods[index+1]
		err := chain.UpdatePodNeighbor(pod, neighbor)
		if err != nil {
			log.WithFields(log.Fields{
				"error":    err,
				"pod":      pod,
				"neighbor": neighbor,
			}).Error("failed updating neighbor")
			continue
		}
	}

	log.WithFields(log.Fields{
		"pods": pods,
	}).Debug("successfull UpdateChainNeighbors")

	w.WriteHeader(http.StatusOK)
}
