package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	redis "../redisInterface"
	"../store"
)

// GetValue ...
func GetValue(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]

	pool := redis.NewPool()
	connection := pool.Get()
	defer connection.Close()

	err := redis.Ping(connection)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed connecting to redis")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	value, err := redis.Get(connection, key)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed getting value")

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

	store := store.Entry{Key: key, Value: value}
	json.NewEncoder(w).Encode(store)
}
