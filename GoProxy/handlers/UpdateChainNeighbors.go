package handlers

import (
	"net/http"

	"github.com/el10savio/Kube-Chain-Replication/GoProxy/chain"

	log "github.com/sirupsen/logrus"
)

// Scatter Gather Goroutine Result
type neighborResult struct {
	pod      string
	neighbor string
	err      error
}

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

	channel := make(chan neighborResult, len(pods)-1)

	// Scatter Goroutines pattern to
	// Send update neighbor request
	// to all pods except the TAIL
	for index := 0; index < cap(channel); index++ {
		pod, neighbor := pods[index], pods[index+1]
		go func() {
			err := chain.UpdatePodNeighbor(pod, neighbor)
			channel <- neighborResult{pod, neighbor, err}
		}()
	}

	// Gather Goroutines
	for index := 0; index < cap(channel); index++ {
		res := <-channel
		if res.err != nil {
			log.WithFields(log.Fields{
				"error":    res.err,
				"pod":      res.pod,
				"neighbor": res.neighbor,
			}).Error("failed updating neighbor")
			continue
		}
	}

	log.WithFields(log.Fields{
		"pods": pods,
	}).Debug("successfull UpdateChainNeighbors")

	w.WriteHeader(http.StatusOK)
}
