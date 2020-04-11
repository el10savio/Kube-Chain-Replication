package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	router "./router"
)

const (
	// PORT ...
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
