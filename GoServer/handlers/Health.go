package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	redis "../redisInterface"
)

// Health ...
func Health(w http.ResponseWriter, r *http.Request) {
	pool := redis.NewPool()
	connection := pool.Get()
	defer connection.Close()

	err := redis.Ping(connection)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed connecting to redis")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Debug("successfull Health")

	w.WriteHeader(http.StatusOK)
}
