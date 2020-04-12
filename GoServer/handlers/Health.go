package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	redis "../redisInterface"
)

// Health is the healthcheck handler to
// check if the service and the
// connection to Redis
// is working
func Health(w http.ResponseWriter, r *http.Request) {
	// Connect to Redis
	pool := redis.NewPool()
	connection := pool.Get()
	defer connection.Close()

	// Check if Redis connection is alright
	err := redis.Ping(connection)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed connecting to redis")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Debug("successfull Health")

	w.WriteHeader(http.StatusOK)
}
