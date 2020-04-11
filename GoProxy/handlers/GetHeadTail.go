package handlers

import (
	"encoding/json"
	"net/http"

	"../chain"

	log "github.com/sirupsen/logrus"
)

// GetHeadTail ...
func GetHeadTail(w http.ResponseWriter, r *http.Request) {
	head, err := chain.GetHead()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed getting head from chain")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tail, err := chain.GetTail()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed getting tail from chain")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	chain := chain.Initializers{HEAD: head, TAIL: tail}

	log.WithFields(log.Fields{
		"chain": chain,
	}).Debug("successfull GetHeadTail")

	json.NewEncoder(w).Encode(chain)
}
