package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	redis "../redisInterface"
	"../store"
)

// GetValue is the handler interface to
// obtain the corresponding value
// from Redis given the key
func GetValue(w http.ResponseWriter, r *http.Request) {
	// Obtain key from URL Params
	key := mux.Vars(r)["key"]

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

	// Get value from Redis
	value, err := redis.Get(connection, key)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed getting value")

		// Value not present 
		// for given key
		if err == redis.ErrNil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"key":   key,
		"value": value,
	}).Debug("successfull GetValue")

	// Return json encoded KV pair 
	store := store.Entry{Key: key, Value: value}
	json.NewEncoder(w).Encode(store)
}
