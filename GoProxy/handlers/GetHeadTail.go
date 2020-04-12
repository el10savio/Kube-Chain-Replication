package handlers

import (
	"encoding/json"
	"net/http"

	"../chain"

	log "github.com/sirupsen/logrus"
)

// GetHeadTail handler returns the HEAD & TAIL of the chain
func GetHeadTail(w http.ResponseWriter, r *http.Request) {
	// Get the HEAD node
	head, err := chain.GetHead()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed getting head from chain")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the TAIL node
	tail, err := chain.GetTail()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed getting tail from chain")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Wrap them in a struct 
	// to json encode them 
	chain := chain.Initializers{HEAD: head, TAIL: tail}

	log.WithFields(log.Fields{
		"chain": chain,
	}).Debug("successfull GetHeadTail")

	json.NewEncoder(w).Encode(chain)
}
