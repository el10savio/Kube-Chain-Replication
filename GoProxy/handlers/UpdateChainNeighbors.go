package handlers

import (
	"net/http"

	"github.com/el10savio/Kube-Chain-Replication/GoProxy/chain"

	log "github.com/sirupsen/logrus"
)

// UpdateChainNeighbors handler updates all
// the nodes in the chain on its 
// respective neighbors except 
// the TAIL node
func UpdateChainNeighbors(w http.ResponseWriter, r *http.Request) {
	// Get active goredis pods
	pods, err := chain.GetPodsList()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed getting pods from cluster")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send update neighbor request to 
	// all pods except the TAIL 
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
