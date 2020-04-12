package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	redis "../redisInterface"
	"../store"
)

const (
	// PORT defines the port value 
	// for the GoServer service
	PORT = ":8080"
)

// SetValue is the handler interface
// to set a key-value in Redis
func SetValue(w http.ResponseWriter, r *http.Request) {
	// json decoded KV pair 
	var store store.Entry
	_ = json.NewDecoder(r.Body).Decode(&store)

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

	// Set key value pair to Redis
	err = redis.Set(connection, store.Key, store.Value)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed setting value")
		return
	}

	// Transmit write pair to chain neighbor
	neighbor := getNeighbor()
	if neighbor != "" {
		err = NeighborSetValue(neighbor, store)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("failed transmitting value to neighbor")
		} else {
			log.WithFields(log.Fields{"neighbor": neighbor}).Debug("successfully transmitted value to neighbor")
		}
	}

	log.WithFields(log.Fields{
		"key":   store.Key,
		"value": store.Value,
	}).Debug("successfull SetValue")

	w.WriteHeader(http.StatusOK)
}

// NeighborSetValue is the handler interface
// to set a key-value in Redis 
// of the neighbor node
func NeighborSetValue(Neighbor string, store store.Entry) error {
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	// json encoded KV pair 
	bytesStore, err := json.Marshal(store)
	if err != nil {
		return err
	}

	// Send POST request to set value in neighbor
	url := "http://" + strings.TrimSpace(Neighbor) + strings.TrimSpace(PORT) + "/store/set"
	_, err = client.Post(url, "application/json", bytes.NewBuffer(bytesStore))
	if err != nil {
		return err
	}

	return nil
}
