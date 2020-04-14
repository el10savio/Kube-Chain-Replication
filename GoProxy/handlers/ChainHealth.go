package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/el10savio/Kube-Chain-Replication/GoProxy/chain"

	log "github.com/sirupsen/logrus"
)

// Scatter Gather Goroutine Result
type healthResult struct {
	status chain.NodeStatus
	err    error
}

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

	channel := make(chan healthResult, len(pods))

	// Scatter Goroutines pattern to
	// Update all pods health status
	for index := 0; index < cap(channel); index++ {
		pod := pods[index]
		go func() {
			status, err := chain.GetPodHealth(pod)
			channel <- healthResult{status, err}
		}()
	}

	health := make([]chain.NodeStatus, 0)

	// Gather Goroutines
	for index := 0; index < cap(channel); index++ {
		res := <-channel
		if res.err != nil || !res.status.Healthy {
			health = append(health, chain.NodeStatus{Node: res.status.Node, Healthy: false})
			continue
		}
		health = append(health, chain.NodeStatus{Node: res.status.Node, Healthy: true})
	}

	log.WithFields(log.Fields{
		"health": health,
	}).Debug("successfull ChainHealth")

	json.NewEncoder(w).Encode(health)
}
