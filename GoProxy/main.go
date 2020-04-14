package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	router "github.com/el10savio/Kube-Chain-Replication/GoProxy/router"
)

const (
	// PORT defines the port value 
	// for the GoProxy service
	PORT = "8090"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	r := router.Router()

	log.WithFields(log.Fields{
		"port": PORT,
	}).Info("starting GoProxy")

	http.ListenAndServe(":"+PORT, r)
}
