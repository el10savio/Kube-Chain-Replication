package handlers

import (
	"encoding/json"
	"net/http"

	"../chain"

	log "github.com/sirupsen/logrus"
)

// ChainHealth gets a healthcheck status 
// of all the nodes in the chain
func ChainHealth(w http.ResponseWriter, r *http.Request) {
	// Get active goredis pods
	pods, err := chain.GetPodsList()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed getting pods from cluster")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	health := make([]chain.NodeStatus, 0)

	// Update all pods health status
	for _, pod := range pods {
		status, err := chain.GetPodHealth(pod)
		if err != nil || !status.Healthy {
			health = append(health, chain.NodeStatus{Node: status.Node, Healthy: false})
			continue
		}
		health = append(health, chain.NodeStatus{Node: status.Node, Healthy: true})
	}

	log.WithFields(log.Fields{
		"health": health,
	}).Debug("successfull ChainHealth")

	json.NewEncoder(w).Encode(health)
}
